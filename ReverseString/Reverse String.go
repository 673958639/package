package main

import (
	"fmt"
)
func main(){
	arr1 := [] byte {'h','e','l','l','o'}
	reverseString(arr1)

	fmt.Println(arr1)
}
func reverseString(s []byte)  {
	for i:=0;i<len(s)/2;i++{
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return
}


