package util

import (
	"crypto/sha256"
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
