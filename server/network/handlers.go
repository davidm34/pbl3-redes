package network

import (
	"fmt"
	"log"
	"net"
	"time"

	"pingpong/server/blockchain"
	"pingpong/server/protocol"
	"pingpong/server/pubsub"
	"pingpong/server/state"
)

// TCPServer lida com todas as interações de rede TCP com os clientes.
type TCPServer struct {
	listenAddr   string
	stateManager *state.StateManager
	broker       *pubsub.Broker
	blockchain   *blockchain.Client // NOVO: Cliente Blockchain injetado
}

// NewTCPServer cria uma nova instância do servidor TCP, injetando as dependências.
// ATUALIZADO: Agora recebe o cliente blockchain
func NewTCPServer(addr string, sm *state.StateManager, broker *pubsub.Broker, bc *blockchain.Client) *TCPServer {
	return &TCPServer{
		listenAddr:   addr,
		stateManager: sm,
		broker:       broker,
		blockchain:   bc,
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
		s.handleOpenPack(player) // Lógica alterada para Blockchain
	case protocol.AUTOPLAY:
		s.handleAutoPlay(player, true)
	case protocol.NOAUTOPLAY:
		s.handleAutoPlay(player, false)
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

func (s *TCPServer) handleOpenPack(player *protocol.PlayerConn) {
	log.Printf("[HANDLER] %s solicitou abrir um pacote. Iniciando transação na Blockchain...", player.ID)

	// 1. Envia a transação para decrementar o stock no contrato inteligente
	txHash, err := s.blockchain.DecrementStock()
	if err != nil {
		log.Printf("[HANDLER] Erro ao enviar transação para %s: %v", player.ID, err)
		s.sendToPlayer(player.ID, protocol.ServerMsg{
			T:    protocol.ERROR,
			Code: protocol.INTERNAL,
			Msg:  "Erro ao comunicar com a Blockchain. Tente novamente.",
		})
		return
	}

	// Notifica o jogador que o processo iniciou (opcional, UX)
	log.Printf("[HANDLER] Transação enviada: %s. Aguardando mineração...", txHash)

	// 2. Aguarda a confirmação da transação (Mineração do bloco)
	// Isso bloqueia a goroutine, mas como cada cliente tem sua conexão ou o handler roda em rotina separada, ok.
	// Idealmente, faríamos isso em background para não travar o processamento de outras mensagens deste jogador.
	go func() {
		err := s.blockchain.WaitForTransactionReceipt(txHash)
		if err != nil {
			log.Printf("[HANDLER] Transação falhou ou timeout para %s: %v", player.ID, err)
			s.sendToPlayer(player.ID, protocol.ServerMsg{
				T:    protocol.ERROR,
				Code: "TX_FAILED",
				Msg:  "A transação na blockchain falhou. O pacote não foi aberto.",
			})
			return
		}

		log.Printf("[HANDLER] Transação confirmada na blockchain! Gerando cartas para %s.", player.ID)

		// 3. Gera as cartas (Stock já foi decrementado na rede)
		// Usamos o PackSystem do StateManager apenas para gerar as cartas, ignorando o stock local dele.
		cards := s.stateManager.PackSystem.GenerateCardsForPack()

		// 4. Lê o stock atualizado da blockchain para exibir
		newStock, _ := s.blockchain.GetStock()

		// 5. Entrega o pacote
		s.sendToPlayer(player.ID, protocol.ServerMsg{
			T:     protocol.PACK_OPENED,
			Cards: cards,
			Stock: int(newStock), // Stock real vindo do contrato
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