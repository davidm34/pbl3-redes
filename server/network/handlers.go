package network

import (
	"fmt"
	"log"
	"net"
	"time"

	"pingpong/server/blockchain"
	"pingpong/server/matchmaking"
	"pingpong/server/protocol"
	"pingpong/server/pubsub"
	"pingpong/server/state"
	"github.com/google/uuid"
)

// TCPServer lida com todas as interações de rede TCP com os clientes.
type TCPServer struct {
	listenAddr   string
	stateManager *state.StateManager
	broker       *pubsub.Broker
	blockchain   *blockchain.Client // NOVO: Cliente Blockchain injetado
	matchmaking  *matchmaking.MatchmakingService
}

// NewTCPServer cria uma nova instância do servidor TCP, injetando as dependências.
// ATUALIZADO: Agora recebe o cliente blockchain
func NewTCPServer(addr string, sm *state.StateManager, broker *pubsub.Broker, bc *blockchain.Client, mm *matchmaking.MatchmakingService) *TCPServer {
	return &TCPServer{
		listenAddr:   addr,
		stateManager: sm,
		broker:       broker,
		blockchain:   bc,
		matchmaking:  mm,
	}
}

// Listen inicia o listener TCP.
func (s *TCPServer) Listen() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return fmt.Errorf("erro ao escutar TCP em %s: %w", s.listenAddr, err)
	}
	defer ln.Close()
	log.Printf("[NETWORK] Servidor TCP a escutar jogadores em %s", s.listenAddr)

	go s.listenForClientMessages()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("[NETWORK] Erro ao aceitar nova conexão: %v", err)
			continue
		}
		go s.handleConn(conn)
	}
}

// listenForClientMessages escuta o tópico do broker e encaminha para o handler apropriado.
func (s *TCPServer) listenForClientMessages() {
	sub := s.broker.Subscribe("client.messages")
	for msg := range sub {
		if action, ok := msg.Payload.(protocol.ClientAction); ok {
			// log.Printf("[HANDLER] <- %s: %s", action.Player.ID, action.Msg.T) // Log verbose opcional
			s.handleMessage(action.Player, action.Msg)
		}
	}
}

// sendToPlayer publica uma mensagem para o tópico de um jogador específico.
func (s *TCPServer) sendToPlayer(playerID string, msg protocol.ServerMsg) {
	s.broker.Publish(fmt.Sprintf("player.%s", playerID), msg)
}

// handleMessage é o roteador de mensagens.
func (s *TCPServer) handleMessage(player *protocol.PlayerConn, msg *protocol.ClientMsg) {
	switch msg.T {
	case protocol.FIND_MATCH:
		s.handleFindMatch(player)
	case protocol.PLAY:
		s.handlePlay(player, msg.CardID)
	case protocol.CHAT:
		s.handleChat(player, msg.Text)
	case protocol.PING:
		s.handlePing(player, msg.TS)
	case protocol.OPEN_PACK:
		s.handleOpenPack(player)
	case protocol.TRADE:        
        s.handleTrade(player, msg.CardID, msg.Text)
	case protocol.GET_COLLECTION:
        s.handleGetCollection(player)
	case protocol.AUTOPLAY:
		s.handleAutoPlay(player, true)
	case protocol.NOAUTOPLAY:
		s.handleAutoPlay(player, false)
	case protocol.MINING_SOLUTION:
		s.handleMiningSolution(player, msg)
	default:
		s.sendToPlayer(player.ID, protocol.ServerMsg{
			T:    protocol.ERROR,
			Code: protocol.INVALID_MESSAGE,
			Msg:  "Tipo de mensagem desconhecido",
		})
	}
}

// --- Handlers Específicos ---

func (s *TCPServer) handleFindMatch(player *protocol.PlayerConn) {
	s.stateManager.AddPlayerToQueue(player)
	s.sendToPlayer(player.ID, protocol.ServerMsg{
		T:    protocol.ERROR, // Nota: O cliente usa ERROR com Code QUEUED para notificação, mantido legado
		Code: "QUEUED",
		Msg:  "Você entrou na fila de matchmaking.",
	})
}

func (s *TCPServer) handlePlay(player *protocol.PlayerConn, cardID string) {
	match := s.stateManager.FindPlayerMatch(player.ID)
	if match == nil {
		s.sendToPlayer(player.ID, protocol.ServerMsg{T: protocol.ERROR, Code: protocol.MATCH_NOT_FOUND, Msg: "Você não está numa partida"})
		return
	}
	if err := match.PlayCard(player.ID, cardID); err != nil {
		s.sendToPlayer(player.ID, protocol.ServerMsg{T: protocol.ERROR, Code: protocol.INVALID_CARD, Msg: err.Error()})
		return
	}
	log.Printf("[HANDLER] Jogada de %s com a carta %s processada.", player.ID, cardID)
}

func (s *TCPServer) handleChat(player *protocol.PlayerConn, text string) {
	match := s.stateManager.FindPlayerMatch(player.ID)
	if match == nil {
		return
	}
	var opponent *protocol.PlayerConn
	if match.P1.ID == player.ID {
		opponent = match.P2
	} else {
		opponent = match.P1
	}
	if opponent != nil {
		s.sendToPlayer(opponent.ID, protocol.ServerMsg{
			T:        protocol.CHAT_MESSAGE,
			SenderID: player.ID,
			Text:     text,
		})
	}
}

func (s *TCPServer) handlePing(player *protocol.PlayerConn, ts int64) {
	now := time.Now().UnixMilli()
	player.LastPing = now
	rtt := now - ts
	s.sendToPlayer(player.ID, protocol.ServerMsg{
		T:     protocol.PONG,
		TS:    ts,
		RTTMs: rtt,
	})
}

func (s *TCPServer) handleMiningSolution(player *protocol.PlayerConn, msg *protocol.ClientMsg) {
	if s.matchmaking == nil {
		s.sendToPlayer(player.ID, protocol.ServerMsg{
			T:    protocol.ERROR,
			Code: protocol.INTERNAL,
			Msg:  "Serviço de matchmaking indisponível.",
		})
		return
	}
	s.matchmaking.SubmitMiningSolution(player, msg.MatchID, msg.Nonce, msg.Hash)
}

// handleOpenPack com Registo de Propriedade (NFT)
func (s *TCPServer) handleOpenPack(player *protocol.PlayerConn) {
	log.Printf("[HANDLER] %s solicitou abrir um pacote.", player.ID)

	// 1. Transação de Débito (Economia)
	txHash, err := s.blockchain.DecrementStock()
	if err != nil {
		s.sendToPlayer(player.ID, protocol.ServerMsg{T: protocol.ERROR, Code: protocol.INTERNAL, Msg: "Erro no Blockchain."})
		return
	}

	go func() {
		// 2. Aguarda confirmação do pagamento
		if err := s.blockchain.WaitForTransactionReceipt(txHash); err != nil {
			s.sendToPlayer(player.ID, protocol.ServerMsg{T: protocol.ERROR, Msg: "Falha no pagamento do pacote."})
			return
		}

		// 3. Gera Cartas Base (Genéricas)
		baseCards := s.stateManager.PackSystem.GenerateCardsForPack()
		
		// 4. Transforma em NFTs (Adiciona UUID único)
		// Ex: "c_001" -> "c_001:550e8400-e29b..."
		uniqueCards := make([]string, len(baseCards))
		for i, baseID := range baseCards {
			uniqueCards[i] = fmt.Sprintf("%s:%s", baseID, uuid.New().String())
		}

		// 5. Regista a Propriedade na Blockchain (AssignCards)
		log.Printf("[HANDLER] A registar posse das cartas %v para %s...", uniqueCards, player.ID)
		assignTx, err := s.blockchain.AssignCards(player.ID, uniqueCards)
		if err != nil {
			log.Printf("[HANDLER] Erro crítico ao atribuir cartas: %v", err)
			// Em produção, teríamos de reembolsar o pacote ou tentar novamente
		} else {
			s.blockchain.WaitForTransactionReceipt(assignTx) // Espera a confirmação da posse
		}

		// 6. Entrega ao Jogador
		newStock, _ := s.blockchain.GetStock()
		s.sendToPlayer(player.ID, protocol.ServerMsg{
			T:     protocol.PACK_OPENED,
			Cards: uniqueCards, // Entrega os IDs únicos!
			Stock: int(newStock),
		})
	}()
}

func (s *TCPServer) handleAutoPlay(player *protocol.PlayerConn, enable bool) {
	player.AutoPlay = enable
	var msg, code string
	if enable {
		code, msg = "AUTOPLAY_ENABLED", "Autoplay ativado."
	} else {
		code, msg = "AUTOPLAY_DISABLED", "Autoplay desativado."
	}
	s.sendToPlayer(player.ID, protocol.ServerMsg{T: protocol.ERROR, Code: code, Msg: msg})
	log.Printf("[HANDLER] %s alterou autoplay para: %t", player.ID, enable)
}

func (s *TCPServer) handleTrade(player *protocol.PlayerConn, cardID string, targetPlayerID string) {
	if cardID == "" || targetPlayerID == "" {
		s.sendToPlayer(player.ID, protocol.ServerMsg{T: protocol.ERROR, Msg: "Uso: /trade <player_id> <card_id>"})
		return
	}

    // Opcional: Verificar se o jogador alvo está online (s.stateManager.IsPlayerOnline...)
    // Mas a blockchain funciona mesmo offline, então vamos deixar passar.

	log.Printf("[HANDLER] %s quer transferir %s para %s", player.ID, cardID, targetPlayerID)

	go func() {
		// Chama o contrato para transferir a posse
		txHash, err := s.blockchain.TransferCard(player.ID, targetPlayerID, cardID)
		
		if err != nil {
			log.Printf("[HANDLER] Erro na transferência: %v", err)
			s.sendToPlayer(player.ID, protocol.ServerMsg{
				T: protocol.ERROR, 
				Msg: "Erro na troca: Você não é o dono desta carta ou erro na rede.",
			})
			return
		}

		// Aguarda confirmação
		if err := s.blockchain.WaitForTransactionReceipt(txHash); err != nil {
			s.sendToPlayer(player.ID, protocol.ServerMsg{T: protocol.ERROR, Msg: "A transação de troca falhou na mineração."})
			return
		}

		// Sucesso! Notifica ambos.
		successMsg := fmt.Sprintf("Sucesso! Carta %s transferida para %s.", cardID, targetPlayerID)
		s.sendToPlayer(player.ID, protocol.ServerMsg{T: protocol.TRADE_SUCCESS, Msg: successMsg})
		
		// Tenta notificar o destinatário se estiver online neste servidor
        // (Num cluster real, usaríamos o broker pub/sub para chegar a outro servidor)
		s.sendToPlayer(targetPlayerID, protocol.ServerMsg{
			T: protocol.TRADE_SUCCESS, 
			Msg: fmt.Sprintf("Você recebeu a carta %s de %s!", cardID, player.ID),
		})
	}()
}

func (s *TCPServer) handleGetCollection(player *protocol.PlayerConn) {
    log.Printf("[HANDLER] %s solicitou visualizar sua coleção.", player.ID)
    
    // Consulta a Blockchain
    cards, err := s.blockchain.GetUserCards(player.ID)
    if err != nil {
        log.Printf("[HANDLER] Erro ao buscar coleção: %v", err)
        s.sendToPlayer(player.ID, protocol.ServerMsg{T: protocol.ERROR, Msg: "Erro ao consultar Blockchain."})
        return
    }

    // Envia resposta
    s.sendToPlayer(player.ID, protocol.ServerMsg{
        T:     protocol.COLLECTION,
        Cards: cards, // Lista de IDs únicos (ex: c_001:uuid...)
    })
}
