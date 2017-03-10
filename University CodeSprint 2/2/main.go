package main

import (
	"fmt"
	"strconv"
)

func checkNumber(num string) string {
	var str1, str2 string
	for k := 1; k <= len(num)/2; k++ {
		str1 = string(num[0:k])
		nn, _ := strconv.Atoi(str1)
		str2 = str1

		for len(str2) < len(num) {
			nn++
			str2 += strconv.Itoa(nn)
		}
		//fmt.Println("str2 ",str2)
		if str2 == num {
			return "YES " + str1
			break
		}
		//fmt.Println("k ",k)
		//fmt.Println("len", len(num)/2)
	}
	//fmt.Println("k ",k)
	return "NO"
}
func main() {
	var n, i int64
	var temp string
	fmt.Scanf("%d\n", &n)

	for i = 1; i <= n; i++ {
		fmt.Scanf("%s\n", &temp)
		fmt.Println(checkNumber(temp));
	}
}
