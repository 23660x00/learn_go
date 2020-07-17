package main

import (
	"os"
	"strings"
    "fmt"
    "time"
)

func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// fmt.Println(s)
}

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func Echo3() {
	strings.Join(os.Args[1:], " ")
	// fmt.Println(strings.Join(os.Args[1:], " "))
}

func TestEcho(f func()) {
    total:=0
    for i:=0;i<100;i++{
        s:=time.Now()
        f()
        d:=time.Since(s)
        total+=int(d)
    }
    fmt.Println(total/100)
}

func main(){
    TestEcho(Echo1)
    TestEcho(Echo2)
    TestEcho(Echo3)
}
