package api

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
)

func hash(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}
func generateID() string {
	const charset = "0123456789"
	rand.NewSource(10)
	id := make([]byte, 5)
	for idx := range id {
		id[idx] = charset[rand.Intn(len(charset))]
	}

	return string(id)
}
