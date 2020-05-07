package main

import (
	"fmt"
	"math"
	"sort"
)
func main(){
	arr := [] int {4,2,1,3}
	fmt.Println(minimumAbsDifference(arr))
}
func minimumAbsDifference(arr []int) [][]int {
	min:=int(math.Pow10(6))
	sort.Ints(arr)
	res:=[][]int{}
	for i:=0;i< len(arr)-1;i++{
		if min>arr[i+1]-arr[i]{
			res=[][]int{{arr[i],arr[i+1]}}
			min=arr[i+1]-arr[i]
		}else if min==arr[i+1]-arr[i]{
			res= append(res, []int{arr[i],arr[i+1]})
		}
	}
	return res
}


