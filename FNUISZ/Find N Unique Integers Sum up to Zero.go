package main

import (
	"fmt"
)
func main(){
	n :=1
	fmt.Println(sumZero(n))
}
func sumZero(n int) []int {
	arr:=[]int {}
	if n==1{
		arr=append(arr,0)
		return arr
	}
	if n%2==0 {
		temp := n / 2
		for i := 0; i <= temp-1; i++ {
			arr = append(arr, -(i + 1))
			arr = append(arr, i+1)
		}
		return arr
	}
	temp := (n-1) / 2
	for i:=0;i<=temp-1;i++ {
		arr=append(arr,-(i+1) )
		arr=append(arr,i+1)
	}
	arr= append(arr, 0)
	return arr
}

