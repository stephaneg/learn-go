package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

//Hash : double hash function
func Hash(s string) []byte {
	var r1, r2 []byte

	// init hash function
	h := sha256.New()

	// first round
	h.Write([]byte(s))
	r1 = h.Sum(nil)

	// second round
	h.Reset()
	h.Write(r1)
	r2 = h.Sum(nil)

	return r2
}

// GetPrivateKey : generate a private key
func GetPrivateKey() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error())
	}
	return privateKey
}

// GetPublicKeyFromPrivateKey : retrieve public key from private key
func GetPublicKeyFromPrivateKey(privateKey *rsa.PrivateKey) *rsa.PublicKey {
	publicKey := &privateKey.PublicKey
	return publicKey
}
