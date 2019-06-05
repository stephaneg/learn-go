package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"util"
)

func main() {

	// starting message

	message := "hello"
	fmt.Println("message : " + message)

	// double has for message
	dh := util.Hash(message)
	fmt.Printf("hash : >%x<\n", dh)

	// generate a oprivate key
	fmt.Println("generating a private key...")
	privateKey := util.GetPrivateKey()
	fmt.Println(privateKey)

	// signature for the original message
	fmt.Println("digital signature for message")
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, dh)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("signature : >%X<\n", signature)

	// verify signature for the original document
	publicKey := util.GetPublicKeyFromPrivateKey(privateKey)
	err2 := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, dh, signature)

	if err2 == nil {
		fmt.Println("signature is ok")
	} else {
		fmt.Println("Signature is KO")
	}

}
