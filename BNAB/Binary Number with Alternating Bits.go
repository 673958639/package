package main

import (
	"fmt"
)
func main(){
	n :=5
	fmt.Println(hasAlternatingBits(n))
}
func hasAlternatingBits(n int) bool {
	t := 2
	for n != 0 {
		if t == n&1 {
			return false
		} else {
			t = n&1
		}
		n >>= 1
	}
	return true
}

