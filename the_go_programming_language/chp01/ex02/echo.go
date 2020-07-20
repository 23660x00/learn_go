package main

import (
    "fmt"
    "os"
)

func main(){
    for i,j := range os.Args[1:]{
        fmt.Printf("%v: %vi\n",i+1,j)
    }
}
