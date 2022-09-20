package password

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var salt string

const currentCost = 14

func init() {
	if strings.HasSuffix(os.Args[0], ".test") {
		salt = "testsalt"
	} else {
		salt = os.Getenv("PASSWORDSALT")
	}
	if salt == "" {
		log.Fatalln("A password salt is required. See https://jsnfwlr.github.io/facemasq/errors/#PasswordSalt for more details")
	}
}

func HashPassword(input string) (hash string, err error) {
	var output []byte
	salted := fmt.Sprintf("%sfaceMasq%s", input, salt)
	obfuscated := fmt.Sprintf("%s", sha256.Sum256([]byte(salted)))
	output, err = bcrypt.GenerateFromPassword([]byte(obfuscated), currentCost)
	hash = string(output)
	return
}

func ConfirmPassword(input, stored string) (newHash string, err error) {
	var cost int
	salted := fmt.Sprintf("%sfaceMasq%s", input, salt)
	obfuscated := fmt.Sprintf("%s", sha256.Sum256([]byte(salted)))
	err = bcrypt.CompareHashAndPassword([]byte(stored), []byte(obfuscated))
	if err != nil {
		return
	}
	cost, err = bcrypt.Cost([]byte(stored))
	if err != nil {
		return
	}
	if cost != currentCost {
		newHash, err = HashPassword(input)
	}
	return
}

func setSalt(input string) {
	salt = input
}
