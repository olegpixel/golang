package main

import (
	"fmt"
)

func gcd(a, b int64) int64 {
	var remainder int64

	for (b > 0) {
		remainder = a % b
		a = b
		b = remainder
	}

	return a
}

func fx(m, n int64) int64 {
	return (m * m + 1) % n
}

func fx2(m, n int64) int64 {
	return (m * m + 3) % n
}

func pf(n int64) int64 {
	var x, y, d, temp int64 = 2, 2, 1, 1

	for d == 1 {
		x = fx(x, n)
		y = fx(fx(y, n), n)
		if x - y  < 0 {
			temp = -1
		} else {
			temp = 1
		}
		d = gcd(temp * (x - y), n)
	}

	if d == n {
		for d == 1 {
			x = fx2(x, n)
			y = fx2(fx(y, n), n)
			if x - y  < 0 {
				temp = -1
			} else {
				temp = 1
			}
			d = gcd(temp * (x - y), n)
		}
	}

	//if d % 3 == 0 {
	//	d = d / 3
	//}

	return d
}

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
	var count, input, div, i, j int64
	var dividers []int64

	fmt.Scanf("%d\n", &count)

	for i=1; i <= count; i++ {
		fmt.Scanf("%d\n", &input)

		for j=2; j <= input/2 && j<=1000000; j++ {
			if input % j == 0 {
				dividers = append(dividers, j)
			}
			for input % j == 0 {
				input /= j
				//fmt.Println(input)
				//fmt.Println(j)
			}
		}
		//if input % 2 == 0 {
		//	dividers = append(dividers, 2)
		//}
		//for input % 2 == 0 {
		//	input /= 2
		//}
		//
		//if input % 3 == 0 {
		//	dividers = append(dividers, 3)
		//}
		//for input % 3 == 0 {
		//	input /= 3
		//}

		for input != 1 {
			div = pf(input)
			dividers = append(dividers, div)
			input /= div
		}

		//fmt.Println(dividers)
		fmt.Println(max(dividers))
		dividers = dividers[:0]
	}
}
