package main

import (
	"fmt"
	"sort"
)


func main() {

	arr:= [] int {0,1,2,3}
	ans:=missingNumber(arr)
	fmt.Println(ans)
}
func missingNumber(nums []int) int{
	sort.Ints(nums)
	value:=0
	if len(nums)==1 {
		if nums[0]==1 {
		return 0
		}
		return 1
	}
	for i:=0;i<len(nums)-1;i++ {
		if nums[i]+1==nums[i+1] {
			if nums[0]!=0 {
				return 0
			}
			value=nums[i+1]+1
			continue
		}
		value=nums[i]+1
		break
	}
	return value
}

