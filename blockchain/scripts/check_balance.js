import hre from "hardhat";

async function main() {
  // Pega a conta configurada no hardhat.config.js
  const [signer] = await hre.ethers.getSigners();
  
  console.log("================ DIAGNÓSTICO ================");
  console.log("1. Endereço que o Hardhat está usando:");
  console.log("   ➡️  " + signer.address);
  
  // Consulta o saldo desse endereço na rede Docker
  const balance = await hre.ethers.provider.getBalance(signer.address);
  console.log("\n2. Saldo deste endereço na rede:");
  console.log("   💰 " + hre.ethers.formatEther(balance) + " ETH");
  
  // Consulta o número do último bloco para ver se estamos sincronizados
  const blockNumber = await hre.ethers.provider.getBlockNumber();
  console.log("\n3. Altura do Bloco Atual (Rede Docker):");
  console.log("   🧱 " + blockNumber);
  console.log("=============================================");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});