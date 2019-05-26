package main

import (
	"fmt"
)

func main() {

	s := "Hello world"
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%16b ", s[i])
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

}
