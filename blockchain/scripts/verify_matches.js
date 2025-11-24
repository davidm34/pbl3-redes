import hre from "hardhat";

async function main() {
  // Endereço do seu contrato (o mesmo do .env)
  const contractAddress = "0x1656F6448f3AF1263893cc724B956f66dc318882"; 

  const PackRegistry = await hre.ethers.getContractFactory("PackRegistry");
  const contract = PackRegistry.attach(contractAddress);

  console.log("🔍 Consultando o Registo de Partidas na Blockchain...");

  // 1. Verifica quantas partidas existem
  const count = await contract.getMatchCount();
  console.log(`📊 Total de partidas registadas: ${count}`);

  // 2. Lista os detalhes de cada uma
  if (count > 0) {
    console.log("\n--- Histórico Imutável ---");
    for (let i = 0; i < count; i++) {
      const match = await contract.matches(i);
      
      // Converte timestamp (BigInt) para data legível
      const date = new Date(Number(match.timestamp) * 1000).toLocaleString();

      console.log(`\n🥊 Partida #${i + 1}`);
      console.log(`   🆔 ID: ${match.matchId}`);
      console.log(`   🏆 Vencedor: ${match.winnerId}`);
      console.log(`   💀 Perdedor: ${match.loserId}`);
      console.log(`   📅 Data: ${date}`);
    }
  } else {
    console.log("⚠️ Nenhuma partida encontrada. Jogue uma partida completa primeiro!");
  }
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});