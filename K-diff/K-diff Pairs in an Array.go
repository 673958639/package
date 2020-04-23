package main

import "fmt"

func main() {
	input := []int{3, 1, 4, 1, 5}
	var answer int
	answer=findPairs(input,2)
	fmt.Println(answer)
}
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	same:=make(map[int]bool)
	diff:=make(map[int]bool)
	for _,num:=range nums {
		if same[num + k] {
			diff[num + k] = true
		}
		if same[num - k] {
			diff[num] = true
		}
		same[num] = true
	}
	return len(diff)
	}
