package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main(){
    wCounts:=make(map[string]int)
    r,err:=os.Open(os.Args[1])
    if err != nil{
        fmt.Fprintf(os.Stderr,"err: %v",err)
    }
    input:=bufio.NewScanner(r)
    input.Split(bufio.ScanWords)

    for input.Scan(){
        w:=strings.ToLower(input.Text())
        wCounts[w]++
    }


    for word,count:=range wCounts{

        fmt.Printf("%s: %d\n",word,count)
    }
}
