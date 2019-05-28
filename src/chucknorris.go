package main

import (
	"fmt"
)

func main() {

	outstr string
	s := "CC"
	fmt.Println("string : " + s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%b ", s[i])
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

}
