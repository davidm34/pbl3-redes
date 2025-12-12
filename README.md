# Attribute War: Blockchain Edition 

Este repositório contém a solução para o **Problema 3** da disciplina **MI - Concorrência e Conectividade (TEC502)**. O projeto consiste em um jogo de cartas multiplayer (1v1) com arquitetura distribuída, onde a posse de ativos (cartas e pacotes) e a validação de transações são garantidas através de uma **Blockchain privada (Ethereum/Geth)** e Smart Contracts.

## 📋 Visão Geral

O **Attribute War** evoluiu de um sistema puramente distribuído para uma aplicação descentralizada (dApp) híbrida. Enquanto a lógica de combate e matchmaking ocorre em servidores Go de alto desempenho, a economia do jogo é auditada em blockchain.

### Principais Funcionalidades

  * **Arquitetura Distribuída:** Cluster de servidores de jogo sincronizados via API REST e Pub/Sub, garantindo tolerância a falhas.
  * **Blockchain Integration:** Registro imutável de abertura de pacotes e posse de cartas utilizando Smart Contracts (`PackRegistry.sol`).
  * **Protocolo Personalizado:** Comunicação Cliente-Servidor via TCP com serialização JSON otimizada.
  * **Matchmaking Distribuído:** Sistema capaz de parear jogadores conectados em servidores distintos (S2S).
  * **Economia de Ativos:** Cartas e pacotes são tratadas como ativos digitais únicos.

-----

## 🏗️ Arquitetura do Sistema

O sistema é composto pelos seguintes contêineres Docker orquestrados:

1.  **Game Servers (x3):** Instâncias replicadas (Go) que gerenciam conexões TCP, lógica de jogo e interagem com a blockchain.
2.  **Blockchain Node (Geth):** Um nó Ethereum privado rodando via Geth, onde o Smart Contract está implantado.
3.  **Client:** Interface de linha de comando (CLI) para interação dos jogadores.

### Diagrama Conceitual

```mermaid
graph TD
    Client[Cliente TCP] -->|JSON| ServerLB[Servidor de Jogo (Go)]
    ServerLB -->|RPC| Geth[Blockchain Node (Geth)]
    ServerLB <-->|REST API| ServerPeer[Outros Servidores]
    Geth -- Smart Contract --> Ledger[(Ledger Imutável)]
```

-----

## 🚀 Como Executar

### Pré-requisitos

  * **Docker** e **Docker Compose** instalados.
  * (Opcional) **Go 1.19+** e **Node.js** para desenvolvimento local.

### 1\. Inicialização Rápida (Docker Compose)

O comando abaixo levanta toda a infraestrutura: 3 servidores de jogo, o nó blockchain e o cliente de teste.

```bash
# Clone o repositório
git clone <url-do-repositorio>
cd pbl3-redes-main

# Subir o ambiente completo (com build forçado)
docker-compose up --build
```

> **Nota:** A primeira execução pode levar alguns minutos para configurar o nó Geth e realizar o deploy dos contratos inteligentes.

### 2\. Acessando os Clientes

Com o cluster rodando, abra novos terminais para simular jogadores:

```bash
# Jogador 1 (Conectado ao Servidor 1)
docker run --rm -it --network pbl3-redes-main_game-net \
  -e SERVER_ADDR=server-1:9000 \
  pingpong-client:latest

# Jogador 2 (Conectado ao Servidor 2 - Teste de Sincronização S2S)
docker run --rm -it --network pbl3-redes-main_game-net \
  -e SERVER_ADDR=server-2:9001 \
  pingpong-client:latest
```

-----

## 🎮 Comandos do Jogo

Dentro do cliente CLI, utilize os seguintes comandos:

| Comando | Descrição |
| :--- | :--- |
| `/find` | Busca um oponente na rede distribuída. |
| `/play <id>` | Joga a carta especificada durante a partida. |
| `/pack` | **Blockchain:** Solicita a abertura de um pacote (transação on-chain). |
| `/hand` | Visualiza as cartas na mão atual. |
| `/quit` | Sai do jogo. |

Consulte as regras completas em [GAME\_RULES.md](https://www.google.com/search?q=GAME_RULES.md).

-----

## 🧪 Testes e Verificação

O projeto inclui suites de testes para validar tanto a concorrência distribuída quanto a integridade da blockchain.

### Testes Automatizados (Go)

```bash
# Entrar na pasta de testes
cd tests

# Teste de concorrência de pacotes (Stress Test)
go test -v stress_packs.go

# Teste de estabilidade do cluster
go test -v stress_cluster_test.go
```

### Scripts de Blockchain (Hardhat)

Para verificar contratos ou saldos diretamente:

```bash
cd blockchain
npx hardhat run scripts/check_balance.js --network localhost
npx hardhat run scripts/verify_matches.js --network localhost
```

-----

## 📂 Estrutura de Pastas

  * `server/`: Código fonte dos servidores de jogo (Go).
      * `blockchain/`: Cliente Go para interação com Ethereum (abigen bindings).
      * `game/`: Lógica de estado e regras da partida.
  * `client/`: Cliente CLI (Go).
  * `blockchain/`: Smart Contracts (Solidity), Genesis block e scripts Hardhat.
  * `docs/`: Documentação da arquitetura e relatórios.
  * `tests/`: Scripts de teste de carga e integração.

-----

## 🛠️ Tecnologias Utilizadas

  * **Linguagem:** Go (Golang)
  * **Blockchain:** Ethereum (Geth Private Net), Solidity
  * **Containerização:** Docker, Docker Compose
  * **Protocolos:** TCP (Game), HTTP (Inter-server), JSON-RPC (Blockchain)

-----

## 👥 Autores

Trabalho desenvolvido para a disciplina de Redes / Sistemas Distribuídos.

  * **Cláudio Daniel Figueredo Peruna**
  * **David Neves Dias**

