package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)
func main(){
	s1 := "abc"
	s2 := "cba"
	fmt.Println(CheckPermutation(s1,s2))
}

func CheckPermutation(s1 string, s2 string) bool {
	arr1:=strings.Split(s1,"")
	arr2:=strings.Split(s2,"")
	sort.Strings(arr1)
	sort.Strings(arr2)
	return reflect.DeepEqual(arr1, arr2)
}
