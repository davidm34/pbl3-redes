package blockchain

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	// Importa o código que geramos com o abigen
	"pingpong/server/blockchain/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client gerencia a comunicação com o Smart Contract
type Client struct {
	ethClient *ethclient.Client
	contract  *contracts.PackRegistry
	auth      *bind.TransactOpts // Autorização (Sua Carteira)
	address   common.Address     // Endereço do Contrato
}

// NewClient inicializa a conexão com o Blockchain com tentativas de reconexão
func NewClient(nodeURL string, contractAddrHex string, privateKeyHex string) (*Client, error) {
	var client *ethclient.Client
	var err error

	// Tenta conectar por 60 segundos (esperando o Geth subir e gerar o DAG)
	log.Printf("[BLOCKCHAIN] Tentando conectar ao nó em %s...", nodeURL)
	for i := 0; i < 12; i++ {
		client, err = ethclient.Dial(nodeURL)
		if err == nil {
			// Tenta fazer uma chamada simples para garantir que o nó está pronto (ex: ChainID)
			_, err = client.ChainID(context.Background())
			if err == nil {
				break // Sucesso total!
			}
		}
		log.Printf("[BLOCKCHAIN] Nó indisponível (%v). Tentando novamente em 5s...", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao nó Ethereum após várias tentativas: %v", err)
	}

	// 2. Configurar a "Carteira" (Signer)
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("chave privada inválida: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("falha ao obter chainID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar transactor: %v", err)
	}

	// 3. Carregar a Instância do Contrato
	address := common.HexToAddress(contractAddrHex)
	instance, err := contracts.NewPackRegistry(address, client)
	if err != nil {
		return nil, fmt.Errorf("falha ao carregar contrato PackRegistry: %v", err)
	}

	log.Printf("[BLOCKCHAIN] Conectado com sucesso! Contrato: %s (ChainID: %v)", contractAddrHex, chainID)

	return &Client{
		ethClient: client,
		contract:  instance,
		auth:      auth,
		address:   address,
	}, nil
}

// GetStock lê o total de pacotes diretamente da Blockchain (Leitura é grátis e instantânea)
func (c *Client) GetStock() (int64, error) {
	// O parâmetro 'nil' usa as opções de chamada padrão (bloco mais recente)
	total, err := c.contract.GetStock(nil)
	if err != nil {
		return 0, fmt.Errorf("erro ao ler estoque do contrato: %w", err)
	}
	return total.Int64(), nil
}

// DecrementStock envia uma transação para diminuir o estoque (Escrita custa Gas e leva tempo)
func (c *Client) DecrementStock() (string, error) {
	// Atualiza o contexto da transação (Gas Price, Nonce, etc)
	// Em redes de produção, precisaríamos calcular o gas price dinamicamente.
	// Na nossa rede privada, o padrão geralmente funciona bem.
	
	// IMPORTANTE: Transações de escrita precisam do 'c.auth' para serem assinadas
	tx, err := c.contract.DecrementStock(c.auth)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar transação: %w", err)
	}

	log.Printf("[BLOCKCHAIN] Transação enviada! Hash: %s. Aguardando confirmação...", tx.Hash().Hex())

	// Opcional: Aguardar a confirmação (mineração) do bloco
	// Para um jogo rápido, talvez não queiramos bloquear aqui, mas para segurança sim.
	// Aqui retornamos o hash imediatamente.
	return tx.Hash().Hex(), nil
}

// WaitForTransactionReceipt aguarda até que a transação seja minerada
func (c *Client) WaitForTransactionReceipt(txHashHex string) error {
	txHash := common.HexToHash(txHashHex)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Aumentei para 30s por segurança
	defer cancel()

	// CORREÇÃO: Primeiro buscamos o objeto da transação na rede
	tx, isPending, err := c.ethClient.TransactionByHash(ctx, txHash)
	if err != nil {
		// Se a transação não for encontrada imediatamente, pode ser que ainda não tenha sido propagada
		return fmt.Errorf("erro ao recuperar a transação %s: %v", txHashHex, err)
	}

	if isPending {
		log.Printf("[BLOCKCHAIN] A transação %s está no pool pendente...", txHashHex)
	}

	// Agora passamos o objeto 'tx' (e não o hash) para o WaitMined
	receipt, err := bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return fmt.Errorf("erro ao aguardar mineração: %w", err)
	}

	// Verificação extra: Status 1 = Sucesso, 0 = Falha
	if receipt.Status == 0 {
		return fmt.Errorf("a transação foi minerada mas FALHOU na execução (revertida)")
	}

	return nil
}