package main

import (
	"fmt"
	"strconv"
)


func solveCase(m, n int32) int64 {
	arr := make([][]int32, 1000)
	friends := make([][]int32, 1000)
	//var arr [][]int32
	//var friends [][]int32
	var temp int64
	var input string
	var i, f1, f2 int32

	for i = 1; i <= m; i++ {
		fmt.Scanf("%s", &input)
		temp, _ = strconv.ParseInt(input, 10, 64);
		f1 = int32(temp)

		fmt.Scanf("%s\n", &input)
		temp, _ = strconv.ParseInt(input, 10, 64);
		f2 = int32(temp)

		arr = append(arr, []int32{f1, f2})
	}

	for i = 1; i <= n; i++ {
		friends[arr[i][0]][0]++
		friends[arr[i][1]][0]++
	}

	for i = 1; i <= m; i++ {

		var max int64
		var index int32
		for i = 1; i <= m; i++ {
			var searchMax int64
			searchMax = int64(friends[arr[i][0]][0]) + int64(friends[arr[i][1]][0])
			if searchMax > max {
				max = searchMax
				index = i
				fmt.Println(index)
			}
		}
		fmt.Println(arr)
		fmt.Println(friends)
		//friends[arr[index][0]][1]
		//friends[arr[index][1]][1]
		//
		//friends[arr[index][0]][2]
		//friends[arr[index][1]][2]
	}
	//fmt.Println(arr)

	return 1

}
func main() {

	var input, longStr string
	//var i, j, k, temp, result, oneDigit, twoDigits, threeDigits int64
	var q, qIndex int8
	var n, m int32
	var temp int64

	fmt.Scanf("%s\n", &input)
	temp, _ = strconv.ParseInt(input, 10, 64);
	q = int8(temp)

	for qIndex = 1; qIndex <= q; qIndex++ {
		fmt.Scanf("%s", &input)
		temp, _ = strconv.ParseInt(input, 10, 64);
		n = int32(temp)

		fmt.Scanf("%s\n", &input)
		temp, _ = strconv.ParseInt(input, 10, 64);
		m = int32(temp)

		fmt.Println (solveCase(m, n))
	}

	fmt.Scanf("%s\n", &longStr)


}