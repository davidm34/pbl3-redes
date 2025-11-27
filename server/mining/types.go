package mining

// MiningChallenge representa o desafio de prova de trabalho enviado aos clientes.
// Difficulty define quantos zeros à esquerda são exigidos no hash em hexadecimal.
// RandomNonce fornece entropia para evitar reuso de desafios e Timestamp ancora o desafio no tempo.
type MiningChallenge struct {
	Difficulty  int
	RandomNonce string
	Timestamp   int64
}

// MiningSolution contém a resposta encontrada pelo minerador para um desafio.
// Nonce é o valor descoberto e Hash é o resultado SHA-256 correspondente.
type MiningSolution struct {
	Nonce uint64
	Hash  string
}
