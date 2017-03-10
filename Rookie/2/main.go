package main

import (
	"fmt"
	"math"
	"sort"
)

type customArr []int64

func (s customArr) Len() int {
	return len(s)
}
func (s customArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s customArr) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	var ts, i, j, temp, minDiff int64
	var numbersArr customArr

	minDiff = math.MaxInt64
	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%d", &temp)
		numbersArr = append(numbersArr, temp)

	}

	sort.Sort(numbersArr)

	for j = 1; j < int64(len(numbersArr)); j++ {
		if int64(math.Abs(float64(numbersArr[j] - numbersArr[j-1]))) < minDiff {
			minDiff = int64(math.Abs(float64(numbersArr[j] - numbersArr[j-1])))
		}
		if minDiff < 1 {
			break
		}
	}

	fmt.Println(minDiff)

}
