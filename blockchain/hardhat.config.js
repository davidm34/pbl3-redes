import "@nomicfoundation/hardhat-toolbox";

// ADICIONE O '0x' NO INÍCIO DA STRING
const PRIVATE_KEY = "0x2c9063953c63132870b25987dd055a15d67c12317f7d6246c5a5071013d3527c";

/** @type import('hardhat/config').HardhatUserConfig */
export default {
  solidity: "0.8.24",
  networks: {
    docker: {
      url: "http://127.0.0.1:8545",
      chainId: 1337,
      accounts: [PRIVATE_KEY]
    }
  }
};