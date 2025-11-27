package mining

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const maxDifficulty = sha256.Size * 2 // Hexadecimal string length for SHA-256

// GenerateChallenge cria um novo desafio de mineração com entropia fresca.
func GenerateChallenge(difficulty int) MiningChallenge {
	normalized := normalizeDifficulty(difficulty)

	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		// Fallback determinístico com base no tempo e dificuldade caso o rand falhe.
		sum := sha256.Sum256([]byte(fmt.Sprintf("%d:%d", time.Now().UTC().UnixNano(), normalized)))
		copy(randomBytes, sum[:len(randomBytes)])
	}

	return MiningChallenge{
		Difficulty:  normalized,
		RandomNonce: hex.EncodeToString(randomBytes),
		Timestamp:   time.Now().UTC().UnixNano(),
	}
}

// ValidateSolution recalcula o hash e garante que possui os zeros exigidos.
func ValidateSolution(challenge MiningChallenge, solution MiningSolution) bool {
	challenge = normalizeChallenge(challenge)
	expectedHash := computeHash(challenge, solution.Nonce)

	if solution.Hash != "" && solution.Hash != expectedHash {
		return false
	}
	return hasRequiredPrefix(expectedHash, challenge.Difficulty)
}

// SolveChallenge encontra um nonce válido em paralelo utilizando todos os núcleos disponíveis.
func SolveChallenge(challenge MiningChallenge) MiningSolution {
	challenge = normalizeChallenge(challenge)
	targetPrefix := prefixForDifficulty(challenge.Difficulty)

	numWorkers := runtime.NumCPU()
	if numWorkers < 1 {
		numWorkers = 1
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultCh := make(chan MiningSolution, 1)
	var wg sync.WaitGroup

	for workerID := 0; workerID < numWorkers; workerID++ {
		wg.Add(1)

		go func(start uint64) {
			defer wg.Done()

			nonce := start
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				hash := computeHash(challenge, nonce)
				if strings.HasPrefix(hash, targetPrefix) {
					select {
					case resultCh <- MiningSolution{Nonce: nonce, Hash: hash}:
						cancel() // Cancela os demais workers assim que encontramos uma solução.
					default:
					}
					return
				}
				nonce += uint64(numWorkers)
			}
		}(uint64(workerID))
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	solution, ok := <-resultCh
	if !ok {
		// Não deveria acontecer, mas evita deadlock em caso de erro inesperado.
		return MiningSolution{}
	}

	// Garante que o Hash está preenchido se algum consumidor ignorar o campo retornado.
	if solution.Hash == "" {
		solution.Hash = computeHash(challenge, solution.Nonce)
	}

	return solution
}

func computeHash(challenge MiningChallenge, nonce uint64) string {
	input := fmt.Sprintf("%d:%s:%d:%d", challenge.Difficulty, challenge.RandomNonce, challenge.Timestamp, nonce)
	sum := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sum[:])
}

func normalizeDifficulty(difficulty int) int {
	switch {
	case difficulty < 0:
		return 0
	case difficulty > maxDifficulty:
		return maxDifficulty
	default:
		return difficulty
	}
}

func prefixForDifficulty(difficulty int) string {
	if difficulty <= 0 {
		return ""
	}
	return strings.Repeat("0", difficulty)
}

func hasRequiredPrefix(hash string, difficulty int) bool {
	return strings.HasPrefix(hash, prefixForDifficulty(difficulty))
}

func normalizeChallenge(challenge MiningChallenge) MiningChallenge {
	challenge.Difficulty = normalizeDifficulty(challenge.Difficulty)
	return challenge
}
