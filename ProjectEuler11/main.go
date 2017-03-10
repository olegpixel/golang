package main

import (
	"fmt"
)


func main() {
	var ts int
	var values [][]int
	var max int64

	for i := 0; i <= 19; i++ {
		var arr []int
		for j := 0; j <= 19; j++ {
			fmt.Scanf("%d", &ts)
			arr = append(arr, ts)
		}
		values = append(values, arr)
	}

	for i := 0; i <= 19; i++ {
		for j := 0; j <= 19; j++ {
			if j >= 3 {
				var temp int64
				temp = int64(values[i][j]) * int64(values[i][j - 1]) * int64(values[i][j - 2]) * int64(values[i][j - 3])
				if temp > max {
					max = temp;
				}
			}
			if i >= 3 {
				var temp int64
				temp = int64(values[i][j]) * int64(values[i-1][j]) * int64(values[i-2][j]) * int64(values[i-3][j])
				if temp > max {
					max = temp;
				}
			}
			if i >= 3 && j >= 3{
				var temp int64
				temp = int64(values[i][j]) * int64(values[i-1][j-1]) * int64(values[i-2][j-2]) * int64(values[i-3][j-3])
				if temp > max {
					max = temp;
				}
			}
			if j >= 3 && i <= 16 {
				var temp int64
				temp = int64(values[i][j]) * int64(values[i + 1][j - 1]) * int64(values[i + 2][j - 2]) * int64(values[i + 3][j - 3])
				if temp > max {
					max = temp;
				}
			}
		}
	}

	fmt.Println(max)
}
