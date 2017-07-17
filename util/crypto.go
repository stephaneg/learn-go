// package util contains utility
package util

import (
	"crypto/sha256"
	"fmt"
)

// double SHA256 hash
func Hash(s string) string {
	var r1, r2 []byte
	var ret string

	// init hash function
	h := sha256.New()

	// first round
	h.Write([]byte(s))
	r1 = h.Sum(nil)
	fmt.Printf(">%x<\n", r1)

	// second round
	h.Reset()
	h.Write(r1)
	r2 = h.Sum(nil)
	fmt.Printf(">%x<\n", r2)
	ret = string(r2)
	return ret

}
