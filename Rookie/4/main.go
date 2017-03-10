package main

import (
	"fmt"
)
type coord struct {
	x int
	y int
	step int64
}
func findPath(n, m, l int) int64 {
	var desk [][]int64
	var stack []coord
	stack = append(stack, coord{0, 0, 0})
	var i, j int

	for i = 1; i <= l; i++ {
		var temp []int64
		for j = 1; j <= l; j++ {
			temp = append(temp, -1)
		}
		desk = append(desk, temp)
	}

	//fmt.Println(desk)
	for i = 0; i < len(stack); i++ {
		var position coord = stack[i]
		desk[position.x][position.y] = position.step

		if (position.y - n >= 0 && position.x - m >= 0) {
			if desk[position.x - m][position.y - n] > position.step + 1 || desk[position.x - m][position.y - n] == -1 {
				desk[position.x - m][position.y - n] = position.step + 1
				stack = append(stack, coord{position.x - m, position.y - n, position.step + 1})
			}
		}
		if (position.y - n >= 0 && position.x + m <= l-1) {
			if desk[position.x + m][position.y - n] > position.step + 1 || desk[position.x + m][position.y - n] == -1 {
				desk[position.x + m][position.y - n] = position.step + 1
				stack = append(stack, coord{position.x + m, position.y - n, position.step + 1})
			}

		}
		if (position.y + n <= l-1 && position.x + m <= l-1) {
			if desk[position.x + m][position.y + n] > position.step + 1 || desk[position.x + m][position.y + n] == -1 {
				desk[position.x + m][position.y + n] = position.step + 1
				stack = append(stack, coord{position.x + m, position.y + n, position.step + 1})
			}
		}
		if (position.y + n <= l-1 && position.x - m >= 0) {
			if desk[position.x - m][position.y + n] > position.step + 1 || desk[position.x - m][position.y + n] == -1 {
				desk[position.x - m][position.y + n] = position.step + 1
				stack = append(stack, coord{position.x - m, position.y + n, position.step + 1})
			}
		}
		if (position.y - m >= 0 && position.x - n >= 0) {
			if desk[position.x - n][position.y - m] > position.step + 1 || desk[position.x - n][position.y - m] == -1 {
				desk[position.x - n][position.y - m] = position.step + 1
				stack = append(stack, coord{position.x - n, position.y - m, position.step + 1})
			}
		}
		if (position.y - m >= 0 && position.x + n <= l-1) {
			if desk[position.x + n][position.y - m] > position.step + 1 || desk[position.x + n][position.y - m] == -1 {
				desk[position.x + n][position.y - m] = position.step + 1
				stack = append(stack, coord{position.x + n, position.y - m, position.step + 1})
			}
		}
		if (position.y + m <= l-1 && position.x + n <= l-1) {
			if desk[position.x + n][position.y + m] > position.step + 1 || desk[position.x + n][position.y + m] == -1 {
				desk[position.x + n][position.y + m] = position.step + 1
				stack = append(stack, coord{position.x + n, position.y + m, position.step + 1})
			}
		}
		if (position.y + m <= l-1 && position.x - n >= 0) {
			if desk[position.x - n][position.y + m] > position.step + 1 || desk[position.x - n][position.y + m] == -1 {
				desk[position.x - n][position.y + m] = position.step + 1
				stack = append(stack, coord{position.x - n, position.y + m, position.step + 1})
			}
		}

	}
	return desk[l-1][l-1]

}

func main() {
	var i, j int
	var ts int

	fmt.Scanf("%d\n", &ts)

	for i = 1; i < ts; i++ {
		for j = 1; j < ts; j++ {
			if j != ts-1 {
				fmt.Print(findPath(i, j, ts)," ")
			} else {
				fmt.Println(findPath(i, j, ts))
			}
		}
	}

}
