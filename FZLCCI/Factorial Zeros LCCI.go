package main

import (
	"fmt"
)
func main(){
	n:=3
	fmt.Println(trailingZeroes(n))
}
func trailingZeroes(n int) int {
	res := 0;
	for n > 0 {
		res += n/5;
		n /= 5;
	}
	return res;
}

