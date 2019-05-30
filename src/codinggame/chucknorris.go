package codinggame

import (
	"bytes"
	"fmt"
)

/*DecodeBinary :  transform a string in a sequence of bits (string)*/
func DecodeBinary(s string) string {
	var buffer bytes.Buffer

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
		fmt.Printf("%b ", s[i])
		fmt.Println()
		buffer.WriteString(fmt.Sprintf("%b", s[i]))
	}
	return buffer.String()
}

func main() {

	var message, binstr, encodedstr string
	var oldchr byte
	var outbuf bytes.Buffer

	message = "CC"
	fmt.Println("string : " + message)

	/*
	   transform message in a string of bits
	*/
	binstr = DecodeBinary(message)
	fmt.Println(binstr)

	/*
	   parse the string of bits
	*/

	oldchr = 0
	for i := 0; i < len(binstr); i++ {

		// nouveau symbole
		if binstr[i] != oldchr {
			oldchr = binstr[i]
			switch oldchr {
			case '0':
				outbuf.WriteString(" 00 0")
			case '1':
				outbuf.WriteString(" 0 0")
			}
		} else {
			outbuf.WriteString("0")
		}
	}

	encodedstr = outbuf.String()
	fmt.Println("encoded : " + encodedstr)
}
