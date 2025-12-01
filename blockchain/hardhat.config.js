import "@nomicfoundation/hardhat-toolbox";


dotenv.config({ path: "../.env" });
// ADICIONE O '0x' NO INÍCIO DA STRING
const PRIVATE_KEY = process.env.ADMIN_PRIVATE_KEY || process.env.PRIVATE_KEY || "";

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