package main

import (
	"log"
	"net/http"
	"os"
	"pingpong/server/api"
	"pingpong/server/blockchain"
	"pingpong/server/matchmaking"
	"pingpong/server/network"
	"pingpong/server/pubsub"
	"pingpong/server/state"
	"pingpong/server/token"
	"strings"
)

func main() {
	log.Println("[MAIN] A iniciar o servidor do jogo...")

	// 1. Configuração da Topologia do Cluster
	tcpAddr := getEnv("LISTEN_ADDR", ":9000")
	apiAddr := getEnv("API_ADDR", ":8000")
	thisServerAddress := "http://" + getEnv("HOSTNAME", "localhost") + apiAddr

	allServersEnv := getEnv("ALL_SERVERS", thisServerAddress)
	allServers := strings.Split(allServersEnv, ",")
	
	// ... (lógica de índice myIndex mantém-se igual) ...
	myIndex := -1
	for i, addr := range allServers {
		if addr == thisServerAddress {
			myIndex = i
			break
		}
	}
	if myIndex == -1 {
		log.Fatalf("[MAIN] Endereço do servidor %s não encontrado na lista ALL_SERVERS", thisServerAddress)
	}

	nextIndex := (myIndex + 1) % len(allServers)
	nextServerAddress := allServers[nextIndex]
	log.Printf("[MAIN] Topologia do anel configurada. Eu sou %s. O próximo é %s.", thisServerAddress, nextServerAddress)

	// 2. Configuração do Blockchain (SEGURA)
	
	blockchainNodeURL := os.Getenv("BLOCKCHAIN_NODE_URL")
	if blockchainNodeURL == "" {
		// Fallback seguro apenas para desenvolvimento local fora do docker, se necessário
		blockchainNodeURL = "http://blockchain-node:8545" 
	}

	contractAddr := os.Getenv("CONTRACT_ADDRESS")
	if contractAddr == "" {
		log.Fatal("[MAIN] Erro fatal: CONTRACT_ADDRESS não configurado no .env")
	}

	adminPrivateKey := os.Getenv("ADMIN_PRIVATE_KEY")
	if adminPrivateKey == "" {
		log.Fatal("[MAIN] Erro fatal: ADMIN_PRIVATE_KEY não configurada no .env")
	}

	log.Printf("[MAIN] A conectar ao Blockchain em %s...", blockchainNodeURL)
	log.Printf("[MAIN] Endereço do Contrato alvo: %s", contractAddr)

	// Inicializa o cliente
	blockchainClient, err := blockchain.NewClient(blockchainNodeURL, contractAddr, adminPrivateKey)
	if err != nil {
		log.Fatalf("[MAIN] Erro fatal ao conectar ao Blockchain: %v", err)
	}
	
	// Testa leitura inicial
	if stock, err := blockchainClient.GetStock(); err == nil {
		log.Printf("[MAIN] Conexão Blockchain OK! Estoque atual de pacotes: %d", stock)
	} else {
		log.Printf("[MAIN] Aviso: Conectado, mas erro ao ler estoque: %v", err)
	}
	
	// 3. Inicialização dos Componentes Centrais
	stateManager := state.NewStateManager(blockchainClient)
	broker := pubsub.NewBroker()
	tokenAcquiredChan := make(chan *token.Token, 1)

	// 4. Serviços
	matchmakingService := matchmaking.NewService(
		stateManager,
		broker,
		tokenAcquiredChan,
		thisServerAddress,
		allServers,
		nextServerAddress,
	)

	apiServer := api.NewServer(
		stateManager,
		broker,
		tokenAcquiredChan,
		thisServerAddress,
	)
	apiServer.SetTokenReceiver(matchmakingService)

	// Servidor TCP
	tcpServer := network.NewTCPServer(
		tcpAddr,
		stateManager,
		broker,
		blockchainClient,
		matchmakingService,
	)

	// 5. Iniciar Goroutines
	go func() {
		log.Printf("[MAIN] A iniciar servidor da API em %s...", apiAddr)
		if err := http.ListenAndServe(apiAddr, apiServer.Router()); err != nil {
			log.Fatalf("[MAIN] Erro fatal no servidor da API: %v", err)
		}
	}()

	go matchmakingService.Run()

	log.Printf("[MAIN] A iniciar servidor TCP para jogadores em %s...", tcpAddr)
	if err := tcpServer.Listen(); err != nil {
		log.Fatalf("[MAIN] Erro fatal no servidor TCP: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
