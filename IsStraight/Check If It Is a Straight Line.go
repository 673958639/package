package main

import (
	"fmt"
)
func main(){
	arr:= [] [] int {{-4,-3},{1,0},{3,-1},{0,-1},{-5,2}}
	fmt.Println(checkStraightLine(arr))
}
func checkStraightLine(coordinates [][]int) bool {
	var k float32
	var temp float32
	for i:=0;i<len(coordinates)-1;i++{
		if i==0 {
			if i+1>len(coordinates)-1 {
				return true
			}
			k= (float32(coordinates[i][1]) - float32(coordinates[i+1][1])) / (float32(coordinates[i][0]) - float32(coordinates[i+1][0]))
			temp=k
		}
		k=(float32(coordinates[i][1]) - float32(coordinates[i+1][1])) / (float32(coordinates[i][0]) - float32(coordinates[i+1][0]))
		if temp!=k {
			return false
		}
	}
	return true
}