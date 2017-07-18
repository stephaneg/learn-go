package main

import (
	"bytes"
	"fmt"
	"learn-go/util"
)

func main() {

	message := "hello"
	dh := util.Hash(message)

	fmt.Println("message : " + message)
	fmt.Printf("hash : >%x<\n", dh)

	for i := range dh {
		fmt.Printf(" %x ", dh[i])
	}
	fmt.Println("temporary byte array : ")

	var tmp = []byte{0x95, 0x95, 0xc9}
	fmt.Printf("temp : "+">%x<\n", tmp)

	sl := dh[:3]
	fmt.Printf("slicing hash : >%x<\n", sl)

	if bytes.Equal(sl, tmp) {
		fmt.Println("SAME")
	}

	// printing the full byte array
	for j := range dh {
		fmt.Printf("@x%x, ", dh[j])
	}

}
