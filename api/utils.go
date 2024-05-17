package api

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/rs/xid"
)

func hash(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}
func generateID() xid.ID {
	guid := xid.New()

	return guid
}
