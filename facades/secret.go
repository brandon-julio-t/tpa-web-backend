package facades

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

var secretSingleton []byte

func init() {
	secretSingleton = make([]byte, 512)
	if _, err := rand.Read(secretSingleton); err != nil {
		log.Fatal(err)
	}
	secretSingleton = []byte(hex.EncodeToString(secretSingleton))
}

func UseSecret() []byte {
	return secretSingleton
}
