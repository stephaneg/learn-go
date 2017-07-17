package main

import (
	"fmt"
	"learn-go/util"
)

func main() {

	message := "hello"
	dh := util.Hash(message)

	fmt.Println(message)
	fmt.Println(dh)

}
