package main

import "fmt"

func main() {

	var ts, i, j, u, k, m uint64

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {

		var numArr []uint64
		var max uint64
		fmt.Scanf("%d\n", &m)

		for j = 1; j <= m; j++ {

			var numArrTemp []uint64
			//fmt.Println("-----------------------------------------------");
			//fmt.Println("j =", j)

			for u = 0; u < j; u++  {

				if u == j - 1 {
					fmt.Scanf("%d\n", &k)
				} else {
					fmt.Scanf("%d", &k)
				}

				if j == 1 {
					numArrTemp = append(numArrTemp, k)
				} else {
					if u == 0 {
						numArrTemp = append(numArrTemp, k + numArr[u])
						//fmt.Println("1 k + numArr[u]", k + numArr[u])
					} else if u >= uint64(len(numArr)) {
						numArrTemp = append(numArrTemp, k + numArr[u - 1])
						//fmt.Println("2 k + numArr[u-1]", k + numArr[u - 1])
					} else if numArr[u] > numArr[u - 1] {
						numArrTemp = append(numArrTemp, k + numArr[u])
						//fmt.Println("3 k + numArr[u]", k + numArr[u])
					} else {
						numArrTemp = append(numArrTemp, k + numArr[u - 1])
						//fmt.Println("4 k + numArr[u-1]", k + numArr[u - 1])
					}
				}
			}
			numArr = numArrTemp
			//fmt.Println("numArr =", numArr)
		}

		for _, e := range numArr {
			if e > max {
				max = e
			}
		}

		fmt.Println(max)
	}
}