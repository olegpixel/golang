package main

import "fmt"

func nod(a, b int64) int64 {
	for a != 0 && b != 0 {
		if (a > b) {
			a = a % b
		} else {
			b = b % a
		}
	}
	return a + b
}

func calc(n int64) int64 {
	var i, result int64
	result = 1
	for i = 2; i <= n; i++ {
		result = result * (i / nod(result, i))
	}
	return result
}

func main() {
	var ts, n, result, i int64

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%d\n", &n)
		result = calc(n)
		fmt.Println(result)
	}
}
