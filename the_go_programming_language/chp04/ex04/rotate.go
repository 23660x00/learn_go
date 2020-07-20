package main

import "fmt"

func rotate(s []int, n int) []int{
    //if > len
    ro:=n%len(s)
    //append to slice
    for _,i:=range s[:ro]{
        s=append(s,i)
    }
    return s[ro:]
}


func main(){
    s:=[]int{1,2,3,4,5,6,7}
    s=rotate(s,2)
    s=rotate(s,2)
    fmt.Print(s)
}
