package main

import (
	"fmt"
)

var array = make([][]map[int64]bool, 1000000)

var result = make([]int64, 10)

func main() {
	var x, y, toll, i, j, k, t, f int64
	var n, e int64

	fmt.Scanf("%d %d\n", &n, &e)

	for i = 1; i <= n; i++ {
		array[i] = make([]map[int64]bool, 1000000)

		for j = 1; j <= n; j++ {
			array[i][j] = make(map[int64]bool, 10)
		}
	}

	for i = 1; i <= e; i++ {
		fmt.Scanf("%d %d %d\n", &x, &y, &toll)
		var rest = toll % 10

		array[x][y][rest] = true
		array[y][x][10-rest] = true
	}

	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			for k = 1; k <= n; k++ {
				if j != k && i != j && i != k {
					for t = 0; t <= 9; t++ {
						for f = 0; f <= 9; f++ {
							if array[j][i][t] && array[i][k][f] {
								var temp = (f + t) % 10
								if (!array[j][k][temp]) {
									i = 1
									k = 1
									j = 1
								}
								array[j][k][temp] = true
								//array[k][j][10 - temp] = true
							}
						}
					}
				}
			}
		}
	}

	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			if i != j {
				for k = 0; k <= 9; k++ {
					if array[i][j][k] {
						result[k]++
					}
				}
			}
		}
	}
	for i = 0; i <= 9; i++ {
		fmt.Println(result[i])
	}
}
