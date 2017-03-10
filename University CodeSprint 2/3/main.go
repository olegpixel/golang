package main

import (
	"fmt"
)

//func calcScore(a, b []int64, x int64) int {
//	//aIndex := len(a) - 1
//	//bIndex := len(b) - 1
//	//
//	//for a[aIndex] + b[bIndex] > x {
//	//	if a[aIndex] > b[bIndex] {
//	//		aIndex--
//	//	} else if a[aIndex] < b[bIndex] {
//	//		bIndex--
//	//	} else if aIndex < bIndex {
//	//		aIndex--
//	//	} else {
//	//		bIndex--
//	//	}
//	//}
//	return 1
//}
func main() {
	var ts, n, m, x, i, j, temp int64
	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%d %d %d\n", &n, &m, &x)
		//var a, b []int64
		//var aSum, bSum int64

		for j = 1; j <= n; j++ {
			if j == n {
				fmt.Scanf("%d\n", &temp)
			} else {
				fmt.Scanf("%d", &temp)
			}
			////if aSum + temp <= x {
			////	aSum += temp
			//	a = append(a, temp)
			////}
		}
		for j = 1; j <= m; j++ {
			if j == m {
				fmt.Scanf("%d\n", &temp)
			} else {
				fmt.Scanf("%d", &temp)
			}
			////if bSum + temp <= x {
			////	bSum += temp
			//	b = append(b, temp)
			////}
		}

		fmt.Println(1)

	}
}
