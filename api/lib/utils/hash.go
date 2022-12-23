package utils

import "crypto/sha256"

func Sha256(message string) (hash string) {
	hasher := sha256.New()
	hasher.Write([]byte(message))

	hash = string(hasher.Sum(nil))
	return
}
