package main

import (
	"fmt"
)

func main() {
	var ts, i, temp, max, maxIndex uint64

	counter := make([]uint64, 6, 6)

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%d", &temp)
		counter[temp]++
	}

	for i = 1; i <= 5; i++ {
		if counter[i] > max {
			maxIndex = i
			max = counter[i]
		}
	}

	fmt.Println(maxIndex)

}
