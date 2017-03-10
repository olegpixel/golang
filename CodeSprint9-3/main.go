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

	var middleTop point = point{queenPoint.x, n + 1}
	var middleBottom point = point{queenPoint.x, 0}

	var middleLeft point = point{0, queenPoint.y}
	var middleRight point = point{n + 1, queenPoint.y}

	var leftTop point
	if (queenPoint.x + queenPoint.y <= n + 1) {
		leftTop = point{0, queenPoint.x + queenPoint.y}
	} else {
		leftTop = point{queenPoint.x - n + queenPoint.y - 1, n + 1}
	}

	var rightTop point
	if (n - queenPoint.x + 1 + queenPoint.y <= n + 1) {
		rightTop = point{n + 1, n + 1 + queenPoint.y - queenPoint.x}
	} else {
		rightTop = point{n + 1 + queenPoint.x  - queenPoint.y, n + 1}
	}

	var leftBottom point
	if (n - queenPoint.x + 1 + queenPoint.y <= n + 1) {
		leftBottom = point{queenPoint.x - queenPoint.y, 0}
	} else {
		leftBottom = point{0, queenPoint.y - queenPoint.x}
	}

	var rightBottom point
	if (queenPoint.x + queenPoint.y <= n + 1) {
		rightBottom = point{queenPoint.x + queenPoint.y, 0}
	} else {
		rightBottom = point{n + 1, queenPoint.y - n + queenPoint.x - 1}
	}

	for i = 1; i <= k; i++ {
		fmt.Scanf("%d %d\n", &x, &y)

		x, y = y, x

		if x == middleTop.x || x == middleBottom.x {
			if y > queenPoint.y && y < middleTop.y {
				middleTop.y = y
			}
			if y < queenPoint.y && y > middleBottom.y {
				middleBottom.y = y
			}
		} else if y == middleLeft.y || y == middleRight.y {
			if x > queenPoint.x && x < middleRight.x {
				middleRight.x = x
			}
			if x < queenPoint.x && x > middleLeft.x {
				middleLeft.x = x
			}
		} else if x - y == rightTop.x - rightTop.y || x - y == leftBottom.x - leftBottom.y {
			if x > queenPoint.x && x < rightTop.x {
				rightTop.x = x
				rightTop.y = y
			}
			if x < queenPoint.x && x > leftBottom.x {
				leftBottom.x = x
				leftBottom.y = y
			}
		} else if n + 1 - x - y == n + 1 - leftTop.x - leftTop.y || n + 1 - x - y == n + 1 - rightBottom.x - rightBottom.y  {
			if x > queenPoint.x && x < rightBottom.x {
				rightBottom.x = x
				rightBottom.y = y
			}
			if x < queenPoint.x && x > leftTop.x {
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
