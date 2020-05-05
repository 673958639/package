package main

import (
	"fmt"
)
func main(){
	fmt.Println(insertBits(0,11111,0,4))
}
func insertBits(N int, M int, i int, j int) int {
	return ((N >> (j + 1)) << (j + 1)) | ((N >> i << i) ^ N)|( M<<i)
}



