package main

import "fmt"
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func findGreater(i ...int) int{
	largest := MinInt

	for _, v := range i {
		if v > largest {
			largest = v
		}
	}

	return largest
}


func main() {

	arr := []int{-1,-2, -33, -55, -12}

	fmt.Println(findGreater(arr...))
	fmt.Println(findGreater(1,2,3))
	fmt.Println(findGreater())

}
