# Versão v1.10.26: A mais estável para redes privadas PoW (com mineração)
FROM ethereum/client-go:v1.10.26

# Copia o arquivo genesis
COPY genesis.json /tmp/genesis.json

# Define o diretório de dados
ENV DATADIR=/root/.ethereum

# Script de inicialização:
# 1. Inicializa o genesis se a pasta 'geth' não existir
# 2. Inicia o nó com as flags de API abertas para o Docker
#    --networkid 1337      : ID da rede (deve bater com o genesis)
#    --mine                : Ativa mineração CPU
#    --miner.threads 1     : Usa 1 thread (funciona na v1.10)
#    --allow-insecure-unlock : Permite desbloquear contas via HTTP (útil para dev)
ENTRYPOINT ["sh", "-c", "if [ ! -d \"$DATADIR/geth\" ]; then geth init /tmp/genesis.json; fi; exec geth --networkid 1337 --http --http.addr 0.0.0.0 --http.port 8545 --http.api eth,net,web3,personal,miner --http.corsdomain '*' --http.vhosts '*' --nat extip:127.0.0.1 --allow-insecure-unlock --mine --miner.threads 1 --miner.etherbase 0x123463a4b065722e99115d6c222f267d9cabb524"]