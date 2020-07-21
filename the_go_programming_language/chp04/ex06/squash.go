package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpace(bs []byte) []byte {
	var isLastSpace bool
	i, j := 0,0

	for j < len(bs) {
		r, s := utf8.DecodeRune(bs[j:])
		if unicode.IsSpace(r) {
			if !isLastSpace {
				bs[i] = ' '
				i++
				isLastSpace = true
			}
		} else {
			isLastSpace = false
			utf8.EncodeRune(bs[i:], r)
			i += s
		}
		j += s
	}
	return bs[:i]

}

func main(){
	x:=`a    bbb     d            sad      d`
	fmt.Println(string(squashSpace([]byte(x))))
}
