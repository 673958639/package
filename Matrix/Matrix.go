package main

import (
	"fmt"
)
func main(){
	matrix := [] [] int {{0,0,0},{0,1,0},{0,0,0}}
	fmt.Println(updateMatrix(matrix))
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0 {
				res[i][j] = 10000
				if i != 0 {
					res[i][j] = min(res[i][j], res[i-1][j]+1)
				}
				if j != 0 {
					res[i][j] = min(res[i][j], res[i][j-1]+1)
				}
			} else {
				res[i][j] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if res[i][j] > 1 {
				if i < m-1 {
					res[i][j] = min(res[i][j], res[i+1][j]+1)
				}
				if j < n-1 {
					res[i][j] = min(res[i][j], res[i][j+1]+1)
				}
			}
		}
	}
	return res
}

