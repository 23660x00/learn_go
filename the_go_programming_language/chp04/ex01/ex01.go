package main

import(
    "fmt"
    "crypto/sha256"
)

func main(){

}

func countDifBits(a,b *[32]byte){
    c:=0
    for i;i<32;i++{
        c+=popCount(a[i],b[i])
    }
    return c
}
