package main

import (
	"fmt"
	"strings"
)
func main(){
	input := "     "
	fmt.Println(replaceSpaces(input,5))
}
func replaceSpaces(S string, length int) string {

	return strings.Replace(S[:length]," ","%20",-1)
}

