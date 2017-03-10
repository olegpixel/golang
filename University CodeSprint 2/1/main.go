package main

import "fmt"

func main() {
	var n, i, temp, max, min, maxCounter, minCounter int64
	fmt.Scanf("%d\n", &n)

	fmt.Scanf("%d", &temp)
	max = temp
	min = temp

	for i = 2; i <= n; i++ {
		fmt.Scanf("%d", &temp)

		if temp > max {
			max = temp
			maxCounter++
		}
		if temp < min {
			min = temp
			minCounter++
		}
	}

	fmt.Println(maxCounter, minCounter)

}
