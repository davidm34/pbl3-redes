package network

import (
	"fmt"
	"log"
	"net"
	"time"
	"pingpong/server/protocol"
	"pingpong/server/pubsub"
)

// handleConn gere o ciclo de vida completo de uma única conexão de cliente.
// É executada numa goroutine para cada cliente.
func (s *TCPServer) handleConn(conn net.Conn) {
	peer := conn.RemoteAddr().String()
	log.Printf("[CLIENT_MGR] Nova conexão TCP de %s", peer)

	// Cria o objeto PlayerConn inicialmente com o IP (provisório)
	player := protocol.NewPlayerConn(peer, conn)

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	
	msg, err := player.ReadMsg()
	if err != nil {
		log.Printf("[CLIENT_MGR] Erro no handshake de login com %s: %v", peer, err)
		conn.Close()
		return
	}

	if msg.T != protocol.LOGIN {
		log.Printf("[CLIENT_MGR] %s enviou %s antes de LOGIN. Desconectando.", peer, msg.T)
		player.SendMsg(protocol.ServerMsg{T: protocol.ERROR, Msg: "Você deve fazer LOGIN primeiro."})
		conn.Close()
		return
	}

	// Atualiza o ID do jogador com o nome enviado
	newID := msg.Text
	if newID == "" {
		newID = peer // Fallback se vier vazio
	}
	player.ID = newID
	
	// Remove o timeout para o funcionamento normal
	conn.SetReadDeadline(time.Time{})
	
	log.Printf("[CLIENT_MGR] Login realizado com sucesso: %s (%s)", player.ID, peer)
	// -----------------------

	// Agora sim, adiciona ao gerenciador de estado com o ID correto
	s.stateManager.AddPlayerOnline(player)

	playerSub := s.broker.Subscribe(fmt.Sprintf("player.%s", player.ID))
	go s.writeLoop(player, playerSub)

	defer func() {
		log.Printf("[CLIENT_MGR] A iniciar limpeza para %s...", player.ID)
		s.broker.Unsubscribe(playerSub)

		if opponent := s.stateManager.CleanupPlayer(player); opponent != nil {
			log.Printf("[CLIENT_MGR] A notificar oponente %s sobre a desconexão.", opponent.ID)
			s.sendToPlayer(opponent.ID, protocol.ServerMsg{
				T:    protocol.ERROR,
				Code: "OPPONENT_DISCONNECTED",
				Msg:  "O seu oponente desconectou-se",
			})
			s.sendToPlayer(opponent.ID, protocol.ServerMsg{
				T:      protocol.MATCH_END,
				Result: protocol.WIN,
			})
		}
		conn.Close()
		log.Printf("[CLIENT_MGR] Conexão com %s encerrada.", player.ID)
	}()

	s.readLoop(player)
}

// readLoop lê continuamente mensagens do socket do cliente e publica-as no broker.
func (s *TCPServer) readLoop(player *protocol.PlayerConn) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[CLIENT_MGR] Panic recuperado em readLoop para o cliente %s: %v", player.ID, r)
		}
	}()

	for {
		msg, err := player.ReadMsg()
		if err != nil {
			log.Printf("[CLIENT_MGR] Erro ao ler do cliente %s: %v. A encerrar a conexão.", player.ID, err)
			return
		}
		if msg == nil {
			log.Printf("[CLIENT_MGR] Cliente %s desconectou-se (EOF). A encerrar a conexão.", player.ID)
			return // Sai do loop, acionando a limpeza.
		}
		s.broker.Publish("client.messages", protocol.ClientAction{Player: player, Msg: msg})
	}
}

// writeLoop escuta um canal de subscrição e escreve as mensagens recebidas para o socket do cliente.
func (s *TCPServer) writeLoop(player *protocol.PlayerConn, sub pubsub.Subscriber) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[CLIENT_MGR] Panic recuperado em writeLoop para o cliente %s: %v", player.ID, r)
		}
	}()

	for msg := range sub {
		if serverMsg, ok := msg.Payload.(protocol.ServerMsg); ok {
			err := player.SendMsg(serverMsg)
			if err != nil {
				log.Printf("[CLIENT_MGR] Erro fatal ao escrever para o cliente %s: %v. A encerrar o write loop.", player.ID, err)
				return
			}
		}
	}
}
