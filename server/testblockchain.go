//go:build ignore

package main

import (
	"fmt"
	"log"
	"pingpong/server/blockchain"
)

func main() {
	// Configurações (Copie os valores que você obteve anteriormente)
	nodeURL := "http://localhost:8545" 
	// Endereço que o script de deploy retornou:
	contractAddr := "0xD3b3f388Cc92868600156fe7881720bB149cE830" 
	// Sua chave privada (a que está no genesis.json/hardhat):
	privateKey := "2c9063953c63132870b25987dd055a15d67c12317f7d6246c5a5071013d3527c"

	fmt.Println("🔄 Iniciando teste de conexão Blockchain...")

	// 1. Conectar
	client, err := blockchain.NewClient(nodeURL, contractAddr, privateKey)
	if err != nil {
		log.Fatalf("❌ Erro fatal: %v", err)
	}
	fmt.Println("✅ Conexão estabelecida!")

	// 2. Ler Estoque
	stock, err := client.GetStock()
	if err != nil {
		log.Fatalf("❌ Erro ao ler estoque: %v", err)
	}
	fmt.Printf("📦 Estoque Inicial no Blockchain: %d pacotes\n", stock)

	// 3. Tentar uma Transação (Comprar Pacote)
	fmt.Println("💸 Tentando decrementar estoque...")
	hash, err := client.DecrementStock()
	if err != nil {
		log.Fatalf("❌ Erro na transação: %v", err)
	}
	
	fmt.Printf("⏳ Transação enviada (%s). Aguardando mineração...\n", hash)
	err = client.WaitForTransactionReceipt(hash)
	if err != nil {
		log.Fatalf("❌ Erro na confirmação: %v", err)
	}
	fmt.Println("✅ Transação confirmada no bloco!")

	// 4. Ler Estoque Novamente
	newStock, err := client.GetStock()
	if err != nil {
		log.Fatalf("❌ Erro ao ler novo estoque: %v", err)
	}
	fmt.Printf("📦 Estoque Final no Blockchain: %d pacotes\n", newStock)
}