package main

import "fmt"

func main() {
	var i, j, result int

	i = 1
	j = 2

	for (i < 4000000 && j < 4000000) {
		if j % 2 == 0 {
			result += j
		}
		tmp := i + j
		i = j
		j = tmp
	}
	fmt.Println(result)
}