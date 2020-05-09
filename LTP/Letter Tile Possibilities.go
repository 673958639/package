package main

import (
	"fmt"
)
func main(){
	title := "AAB"
	fmt.Println(numTilePossibilities(title))
}

func numTilePossibilities(tiles string) int {
	used := make([]bool, len(tiles))
	m := make(map[string]struct{})
	backtrack("", used, tiles, m)
	return len(m)
}

func backtrack(now string, used []bool, str string, m map[string]struct{}) {
	if len(now) > 0 {
		m[now] = struct{}{}
	}
	if len(now) == len(str) {
		return
	}
	for i := 0; i < len(str); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		backtrack(now + string(str[i]), used, str, m)
		used[i] = false
	}
}


// 链接：https://leetcode-cn.com/problems/letter-tile-possibilities/solution/golang-36ms-bao-li-hui-su-by-resara/
// 来源：力扣（LeetCode）
