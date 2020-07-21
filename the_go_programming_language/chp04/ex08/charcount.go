package main

import (
    "os"
    "fmt"
    "unicode"
    "unicode/utf8"
    "bufio"
    "io"
)

func main(){

    counts=make(map[string]int)
    invalid := 0

    input := bufio.NewReader(os.Stdin)
    for {
        r,n,err:=input.ReadRune()
        if err == io.EOF{
            break
        }
        if err!=nil{
            fmt.Fprintf(os.Stderr,"charcount: %v\n", err)
            os.Exit(1)
        }
        if r==unicode.ReplacementChar && n == 1 {
            invalid++
            continue
        }
        if unicode.IsDigit(r){
            counts["digit"]++
        }

        //if so on...
    }

    for k,v := range counts{
        fmt.Printf("%s: %d",k,v)
    }


    fmt.Printf("invalid: ",invalid)

}
