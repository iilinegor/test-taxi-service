package storage

import "math/rand"

var (
	alphabet = []rune("abcdefghijklmnopqrstuvwxyz")
)

// getUniqueToken generates 2 letter token
func getUniqueToken() string {
	b := make([]rune, TokenLength)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}
