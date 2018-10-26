package addit

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func createHash(plain string) (string) {
	enc := sha256.New()
	enc.Write([]byte(plain))
	hashbyte := enc.Sum(nil)

	return fmt.Sprintf("%X", hashbyte)
}
