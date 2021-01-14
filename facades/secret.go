package facades

import (
	"crypto/rand"
	"log"
)

var secretSingleton []byte

func init() {
	secretSingleton = make([]byte, 512)
	if _, err := rand.Read(secretSingleton); err != nil {
		log.Fatal(err)
	}
}

func UseSecret() []byte {
	return secretSingleton
}
