package main

import (
	"fmt"
)

func main() {
	var left int
	var right int
	var answer [] int
	fmt.Println("Please enter left:")
	fmt.Scan(&left)
	fmt.Println("Please enter right:")
	fmt.Scan(&right)
	answer=selfDividingNumbers(left,right)
	fmt.Println(answer)
}
func selfDividingNumbers(left int, right int) []int {

	answer:=[]int {}
	for i:=left;i<=right;i++ {
		if check(i) {
			answer=append(answer,i)
		}
	}
	return answer
}
func check(num int) bool{
	tem:=num
	for tem>0 {
		d:=tem%10
		if d==0 {
			return false
		}
		r:=num%d
		if r!=0 {
			return false
		}
		tem/=10
	}
	return true
}