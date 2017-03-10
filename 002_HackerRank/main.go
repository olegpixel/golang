package main

import (
	"fmt"
	"strconv"
)
func main() {

	var result int64
	var text string
	var resultArr = make([]int64,0)

	fmt.Scanf("%s\n", &text)
	count, _ := strconv.ParseInt(text, 10, 64);

	for i:=1; int64(i) <= count; i++ {

		fmt.Scanf("%s\n", &text)
		val, _ := strconv.ParseInt(text, 10, 64)
		result = 0;

		result = (val*(val*val-1)*(3*val+2))/12;

		resultArr = append(resultArr,result)

	}
	for _, k := range resultArr {
		fmt.Println(k)
	}
}
