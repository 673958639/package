package main

import (
	"fmt"
	"sort"
)
func main(){
	arr := [] int {0,0,1,0,0}
	fmt.Println(isStraight(arr))
}
func isStraight(nums []int) bool {
	sort.Ints(nums)
	count := 0
	for i := 0; i < 5; i++ {
		if nums[i] == 0 {
			count = count + 1
		}
	}
	if count == 5|4 {
		return true
	}
	if count == 3 {
		for i := 4; i < len(nums); i++ {
			if nums[i]-nums[i-1] > 4 || nums[i] == nums[i-1] {
				return false
			}
		}
	}

	if count == 2 {
		for i := 3; i < len(nums); i++ {
			if nums[i]-nums[i-1] > 3 || nums[i] == nums[i-1] {
				return false
			}
		}
	}

	// 0 1 2 5 6
	if count == 1 {
		for i := 2; i < len(nums); i++ {
			if nums[i]-nums[i-1] > 2 || nums[i] == nums[i-1] {
				return false
			}
		}

	}
	if count == 0 {
		for i := 1; i < len(nums); i++ {
			if nums[i]-nums[i-1] > 1 || nums[i] == nums[i-1] {
				return false
			}
		}
	}
	return true

}


