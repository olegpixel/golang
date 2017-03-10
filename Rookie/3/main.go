package main

import (
	"fmt"
)

const hackStr = "hackerrank"

func checkStr(str string) string {
	var i, j int64

	for j = 0; j < int64(len(str)); j++ {
		if str[j] == hackStr[i] {
			i++
			if i >= int64(len(hackStr)) {
				return "YES"
				break
			}
		}
	}
	//fmt.Println(i)
	return "NO"

}

func main() {
	var ts, i int64
	var inputStr string

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%s\n", &inputStr)
		fmt.Println(checkStr(inputStr))
	}

}
