package main

import "fmt"

func calc(n int) int {
	var max = 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			var mult = i * j
			if mult > 100000 && mult < n {
				var temp = mult % 1000
				temp = temp % 10 * 100 + temp / 10 % 10 * 10 + temp / 100
				if temp == mult / 1000 {
					if mult > max {
						max = mult
					}
				}
			}
		}
	}
	return max
}

func main() {
	var ts, n, result int

	fmt.Scanf("%d\n", &ts)

	for i := 1; i <= ts; i++ {
		fmt.Scanf("%d\n", &n)
		result = calc(n)
		fmt.Println(result)
	}
}
