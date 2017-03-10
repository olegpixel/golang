package main

import (
	"fmt"
)

func triplet(n int64) int64 {
	var a, b, c int64
	var max int64

	if (n % 2 == 0) {
		for a = 1; a < int64(n/3); a++ {
			b = n * (n - 2 * a) / 2 / (n - a)
			//for b = a+1; b < n-2; b++ {
				c = n - a - b
				if a*a + b*b == c*c && a*b*c > max {
					max = a*b*c
					//fmt.Println(max)
				}
			//}
		}
	}
	if max > 0 {
		return max
	} else {
		return -1
	}
}

func main() {
	var ts, i, n int64

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%d\n", &n)
		fmt.Println(triplet(n))

	}

}
