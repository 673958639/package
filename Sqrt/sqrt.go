package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	var sqrt int
	sqrt=mySqrt(x)
	fmt.Println("Input:",x," Output:",sqrt)
	}
func mySqrt(x int) int {
	if(x<1){
		return x
	}
	start:=0
	end:=x/2+1
	var answer int
	for start=0;start<=end;start++ {
		if start*start==x {
			answer=start
		}else if start*start>x{
			answer=start-1
			break	//When start is greater than x, assign the value of start -1 to the answer
		}
	}
	return answer
}
