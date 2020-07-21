package main

import(
    "fmt"
    "unicode/utf8"
)

func reverse(b []byte) []byte{
    for i,j:=0,len(b)-1;i<j;i,j=i+1,j-1{
        b[i],b[j]=b[j],b[i]
    }

    return b
}

func rUTF8(b []byte) []byte{
    for i:=0;i<len(b);{
        _,s:=utf8.DecodeRune(b[i:])
        reverse(b[i:i+s])
        i+=s
    }

    reverse(b)
    return b
}

func main(){
    s:="我吃饱了"
    fmt.Println(string(rUTF8([]byte(s))))
}
