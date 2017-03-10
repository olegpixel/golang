package main

import (
	"fmt"
	"math"
)

func max(arr []int64) int64 {
	var result int64
	for i:=0; i < len(arr); i++ {
		if arr[i] > result {
			result = arr[i]
		}
	}
	return result
}

func main() {
	var count, input, i, j int64
	var dividers []int64

	fmt.Scanf("%d\n", &count)

	for i=1; i <= count; i++ {
		fmt.Scanf("%d\n", &input)

		for input % 2 == 0 {
			input /= 2
			dividers = append(dividers, 2)
		}

		for j=3; input != 1 && j < int64(math.Sqrt(float64(input))) + 1; j = j+2 {

			for input % j == 0 {
				input /= j
				dividers = append(dividers, j)
			}
		}
		dividers = append(dividers, input)
		fmt.Println(max(dividers))
		dividers = dividers[:0]
	}
	//fmt.Println(dividers)
	//fmt.Println(max(dividers))
}
