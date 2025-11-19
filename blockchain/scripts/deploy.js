import hre from "hardhat";

async function main() {
  console.log("A iniciar o deploy do PackRegistry...");

  // 1. Compila e prepara o contrato
  const PackRegistry = await hre.ethers.getContractFactory("PackRegistry");

  // 2. Envia a transação de deploy
  // Estoque inicial de 1000 pacotes
  const packRegistry = await PackRegistry.deploy(1000);

  console.log("Transação enviada. A aguardar confirmação...");

  // 3. Aguarda a confirmação
  await packRegistry.waitForDeployment();

  const address = await packRegistry.getAddress();
  console.log(`✅ Sucesso! PackRegistry deployado no endereço: ${address}`);
  console.log("GUARDE ESTE ENDEREÇO! O servidor Go precisará dele.");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});