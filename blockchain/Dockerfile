FROM ethereum/client-go:v1.10.26

COPY genesis.json /tmp/genesis.json

ENV DATADIR=/root/.ethereum

# Atualizado: --miner.etherbase agora aponta para 0x3B9AeE...
ENTRYPOINT ["sh", "-c", "if [ ! -d \"$DATADIR/geth\" ]; then geth init /tmp/genesis.json; fi; exec geth --networkid 1337 --http --http.addr 0.0.0.0 --http.port 8545 --http.api eth,net,web3,personal,miner --http.corsdomain '*' --http.vhosts '*' --nat extip:127.0.0.1 --allow-insecure-unlock --mine --miner.threads 1 --miner.etherbase 0x3B9AeE166F388d81Df4950C213289b03769C9799"]