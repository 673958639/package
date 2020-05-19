func pivotIndex(nums []int) int {
	sum:=0
	for i:=0;i< len(nums);i++{
		sum+=nums[i]
	}
	for i:=0;i< len(nums);i++{
		leftSum:=0
		for j:=0;j<i;j++{
			leftSum+=nums[j]
		}
		if leftSum==sum-leftSum-nums[i]{
			return i
		}
	}
	return -1
}
