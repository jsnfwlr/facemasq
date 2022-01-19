package password

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type param struct {
	Password string
	Hash     string
	NewHash  bool
	Error    bool
	ErrorMsg string
}

var Params = []param{
	{Password: "test1", Hash: "$2a$14$Tz3WiOVCIIwxkIktZgRZauZ4jrYM7NYa20z1wipv9WpBHPs0G1C.6", NewHash: false, Error: false, ErrorMsg: ""},                                                                                                                                                                                                 // pass check
	{Password: "test2", Hash: "$2a$10$UFgyWHLQDHxPugEl.aPC2ehCKUDcnG55a9gxQvmWA14wnG0c0dszS", NewHash: true, Error: false, ErrorMsg: ""},                                                                                                                                                                                                  // pass check, but should increase cost of hash and change $2a$10$ prefix to $2a$14$
	{Password: "test3", Hash: "$2a$14$QR/3vMuULNW2puz.2fS2Cu0InHUEDg6JXJ5GOAXjYCPVim8AfiVMx", NewHash: false, Error: true, ErrorMsg: "crypto/bcrypt: hashedPassword is not the hash of the given password"},                                                                                                                               // should throw MismatchedHashAndPassword error
	{Password: "test4", Hash: "$3a$14$0MLxR02OJHMTCovp3epDIuCD9R5W72d8JHBbqD5NgeQsddwA1NbjO", NewHash: false, Error: true, ErrorMsg: "crypto/bcrypt: bcrypt algorithm version '3' requested is newer than current version '2'"},                                                                                                           // should throw HashVersionTooNew error
	{Password: "test5", Hash: "@2a@14@ztdeq7HxPcsB1wfNE.p2wOFf96OYWyx0X/TIRriBiKZz4bMoSIXQO", NewHash: false, Error: true, ErrorMsg: "crypto/bcrypt: bcrypt hashes must start with '$', but hashedSecret started with '@'"},                                                                                                               // should throw InvalidPrefix error
	{Password: "test6", Hash: "kmkbZ884YxlA4kLSbvkqF.wu3iLKKvfzU.Gt84iT3p6y42tVljDCy", NewHash: false, Error: true, ErrorMsg: "crypto/bcrypt: hashedSecret too short to be a bcrypted password"},                                                                                                                                          // should throw HashTooShort error
	{Password: "test7", Hash: fmt.Sprintf("%s%d%s", "$2a$", (bcrypt.MaxCost + 2), "$vY99BpU8qcfV9tdpXUEukerx9fyjogK5pMS62fRszsjHFlafWilVa"), NewHash: false, Error: true, ErrorMsg: fmt.Sprintf("%s%d%s%d%s%d%s", "crypto/bcrypt: cost ", (bcrypt.MaxCost + 2), " is outside allowed range (", bcrypt.MinCost, ",", bcrypt.MaxCost, ")")}, // should throw InvalidCost error
	{Password: "test8", Hash: "$2a$18$E4eMx6vZZD/6ERt7i0oWte7HF0UDmss.ZlTBc3gO8qF3IDaLMJ5.2", NewHash: true, Error: false, ErrorMsg: ""},                                                                                                                                                                                                  // pass check, but should decrease cost of hash and change $2a$18$ prefix to $2a$14$
}

func TestCreatePassword(test *testing.T) {
	setSalt("AlphabetSoup")
	for _, param := range Params {
		_, err := HashPassword(param.Password)
		if err != nil && !param.Error {
			test.Errorf("%v", err)
		}
	}

}

func TestCheckPassword(test *testing.T) {
	setSalt("AlphabetSoup")
	for _, param := range Params {
		hash, err := ConfirmPassword(param.Password, param.Hash)
		if err != nil && !param.Error {
			test.Errorf("%v", err)
		} else if err == nil && param.Error {
			test.Error("The provided password should not match the provided hash")
		} else if err != nil && param.Error && param.ErrorMsg != err.Error() {
			test.Errorf("Expected error:\n\t\t%v\n\tgot:\n\t\t%v\n", param.ErrorMsg, err.Error())

		}
		if hash == "" && param.NewHash {
			test.Error("The provided password should have required a new hash")
		} else if hash != "" && !param.NewHash {
			test.Error("The provided password should not have required a new hash")
		}
		if hash != "" && param.NewHash {
			cost, err := bcrypt.Cost([]byte(hash))
			if err != nil {
				test.Error(err.Error())
			}
			if cost != currentCost {
				test.Errorf("The updated hash has the wrong cost: %d should be %d", cost, currentCost)
			}
		}
	}
}
