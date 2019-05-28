package main

import (
	"fmt"
)

func main() {

	var strs []string

	s := "CC"
	fmt.Println("string : " + s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
		fmt.Println()
		fmt.Printf("%b ", s[i])
		fmt.Println()
		strs = append(strs, fmt.Sprintf("%b ", s[i]))

	}

}
