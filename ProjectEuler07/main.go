package main

import (
	"fmt"
	"strconv"
)

var result [10]int

func stack(m int) {

	for i := 0; i < 9; i++ {
		result[i] = result[i+1]
	}
	result[9] = m

}

func main() {
	var ts, i, j int64
	var s []string
	var input string

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%s\n", &input)
		s = append(s, input)
	}

	var temp = 0;
	for i = 49; i >= 0; i-- {
		for j = 0; j < ts; j++ {
			mtemp, _ := strconv.Atoi(string(s[j][i]))
			temp += mtemp
		}
		stack(int(temp % 10))
		temp /= 10
	}
	for temp > 0 {
		stack(int(temp % 10))
		temp /= 10
	}

	for i = 9; i >= 0; i-- {
		fmt.Printf("%v", result[i])
	}
	//fmt.Println(result)
}
