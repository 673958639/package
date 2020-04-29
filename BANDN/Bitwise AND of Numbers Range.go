package main

import (
	"fmt"
)
func main(){
	m := 5
	n := 7
	fmt.Println(rangeBitwiseAnd(m,n))
}
func rangeBitwiseAnd(m int, n int) int {
	t := 0
	for ; m != n; t++ {
		m >>= 1
		n >>= 1
	}
	return n << t
}



