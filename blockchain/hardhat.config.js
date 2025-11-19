require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.24",
  networks: {
    // Esta rede 'docker' conecta ao seu container Geth rodando localmente
    docker: {
      url: "http://127.0.0.1:8545",
      chainId: 1337,
      // A conta com saldo definida no genesis.json
      accounts: ["0x2c9063953c63132870b25987dd055a15d67c12317f7d6246c5a5071013d3527c"] 
    }
  }
};