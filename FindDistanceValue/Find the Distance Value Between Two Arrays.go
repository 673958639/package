package main

import (
	"fmt"
)
func main(){
	arr1 := [] int {4,-3,-7,0,-10}
	arr2 := [] int {10}
	d := 69
	fmt.Println(findTheDistanceValue(arr1,arr2,d))
}
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for i:=0;i<=len(arr1)-1;i++ {
		count := 0
		for j:=0;j<=len(arr2)-1;j++{
			if abs(arr1[i]-arr2[j])<=d {
				count=0
				break
			}
			count++
		}
		if count==len(arr2) {
			res++
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}



