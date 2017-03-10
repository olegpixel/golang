package main

import (
	"fmt"
)

type point struct {
	x int64
	y int64
}
func main() {
	var x, y, i int64
	var n, k, xQueen, yQueen int64
	var queenPoint point


	fmt.Scanf("%d %d\n", &n, &k)
	fmt.Scanf("%d %d\n", &xQueen, &yQueen)

	xQueen, yQueen = yQueen, xQueen

	queenPoint = point{xQueen, yQueen}

	var leftTop point = point{0, queenPoint.y + queenPoint.x}
	var rightTop point = point{n + 1, queenPoint.y + (n - queenPoint.x) + 1}

	var middleTop point = point{queenPoint.x, n + 1}
	var middleBottom point = point{queenPoint.x, 0}

	var middleLeft point = point{0, queenPoint.y}
	var middleRight point = point{n + 1, queenPoint.y}

	var leftBottom point = point{0, queenPoint.y - queenPoint.x}
	var rightBottom point = point{n + 1, queenPoint.y - (n - queenPoint.x) - 1}

	//fmt.Println("leftTop ", leftTop)
	//fmt.Println("rightTop ", rightTop)
	//fmt.Println("middleTop ", middleTop)
	//fmt.Println("middleBottom ", middleBottom)
	//
	//fmt.Println("middleLeft ", middleLeft)
	//fmt.Println("middleRight ", middleRight)
	//fmt.Println("leftBottom ", leftBottom)
	//fmt.Println("rightBottom ", rightBottom)

	for (leftTop.y > n + 1) {
		leftTop.y--
		leftTop.x++
	}

	for (rightTop.y > n + 1) {
		rightTop.y--
		rightTop.x--
	}

	for (leftBottom.y < 0) {
		leftBottom.y++
		leftBottom.x++
	}

	for (rightBottom.y < 0) {
		leftBottom.y++
		leftBottom.x--
	}

	//fmt.Println("leftTop ", leftTop)
	//fmt.Println("rightTop ", rightTop)
	//fmt.Println("middleTop ", middleTop)
	//fmt.Println("middleBottom ", middleBottom)
	//
	//fmt.Println("middleLeft ", middleLeft)
	//fmt.Println("middleRight ", middleRight)
	//fmt.Println("leftBottom ", leftBottom)
	//fmt.Println("rightBottom ", rightBottom)

	for i = 1; i <= k; i++ {
		fmt.Scanf("%d %d\n", &x, &y)

		x, y = y, x

		if x == middleTop.x || x == middleBottom.x {
			if y > queenPoint.y {
				middleTop.y = y
			}
			if y < queenPoint.y {
				middleBottom.y = y
			}
		} else if y == middleLeft.y || x == middleRight.y {
			if x > queenPoint.x {
				middleRight.x = x
			}
			if x < queenPoint.x {
				middleLeft.x = x
			}
		} else if x - y == rightTop.x - rightTop.y || x - y == leftBottom.x - leftBottom.y {
			if x > queenPoint.x {
				rightTop.x = x
				rightTop.y = y
			}
			if x < queenPoint.x {
				leftBottom.x = x
				leftBottom.y = y
			}
		} else if x + y == leftTop.x + leftTop.y || x + y == rightBottom.x + rightBottom.y {
			if x > queenPoint.x {
				rightBottom.x = x
				rightBottom.y = y
			}
			if x < queenPoint.x {
				leftTop.x = x
				leftTop.y = y
			}
		}
	}

	//fmt.Println("leftTop ", leftTop)
	//fmt.Println("rightTop ", rightTop)
	//fmt.Println("middleTop ", middleTop)
	//fmt.Println("middleBottom ", middleBottom)
	//
	//fmt.Println("middleLeft ", middleLeft)
	//fmt.Println("middleRight ", middleRight)
	//fmt.Println("leftBottom ", leftBottom)
	//fmt.Println("rightBottom ", rightBottom)

	var result int64
	result += middleRight.x - queenPoint.x - 1
	result += queenPoint.x - middleLeft.x - 1
	result += middleTop.y - queenPoint.y - 1
	result += queenPoint.y - middleBottom.y - 1
	result += rightTop.x - queenPoint.x - 1
	result += rightBottom.x - queenPoint.x - 1
	result += queenPoint.x - leftBottom.x - 1
	result += queenPoint.x - leftTop.x - 1
	fmt.Println(result)
}
