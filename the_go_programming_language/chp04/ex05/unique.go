package main

import (
	"fmt"
)

func unique(strs []string) []string {

	i,j:=0,1
	for j<len(strs){
		if strs[i]==strs[j]{
			j++
		}else{
			i++
			strs[i]=strs[j]
		}
	}
	return strs[:i+1]
}

func main() {
	s:=[]string{"a","a","a","b","b","c","c"}
	s=unique(s)
	fmt.Printf("%v",s)
}
