package main

import (
	"fmt"
	"strconv"
)

func greatestProduct(n string, k int64) int64 {
	var i, j int64
	var max int64

	for i = 0; i < int64(len(n)) - k; i++ {
		var temp int64 = 1
		for j = 0; j < k; j++ {
			temp2, _ := strconv.ParseInt(string(n[i+j]), 10, 64)
			temp *= int64(temp2)
		}
		if temp > max {
			max = temp
		}
	}

	return max
}

func main() {
	var ts, i, k int64
	var n, d string

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%s %d\n", &d, &k)
		fmt.Scanf("%s\n", &n)
		fmt.Println(greatestProduct(n, k))

	}

}
