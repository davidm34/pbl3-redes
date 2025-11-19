import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";

// Configuração de caminhos para ESM
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

async function main() {
  // Caminho onde o Hardhat guardou a compilação
  const artifactPath = path.join(
    __dirname,
    "../artifacts/contracts/PackRegistry.sol/PackRegistry.json"
  );

  if (!fs.existsSync(artifactPath)) {
    console.error("❌ Erro: Artifact não encontrado. Execute 'npx hardhat compile' primeiro.");
    process.exit(1);
  }

  const artifact = JSON.parse(fs.readFileSync(artifactPath, "utf8"));

  // 1. Salva o ABI
  const abiPath = path.join(__dirname, "../PackRegistry.abi");
  fs.writeFileSync(abiPath, JSON.stringify(artifact.abi));
  console.log(`✅ ABI extraído para: ${abiPath}`);

  // 2. Salva o Bytecode (Bin) - Importante: Removemos o '0x' inicial
  const binPath = path.join(__dirname, "../PackRegistry.bin");
  fs.writeFileSync(binPath, artifact.bytecode.replace(/^0x/, "")); 
  console.log(`✅ Bytecode extraído para: ${binPath}`);
}

main();