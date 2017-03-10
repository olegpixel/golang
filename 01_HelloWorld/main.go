package main

import (
	"fmt"
	"sort"
)

type num struct {
	n int64
	index int64
	res int64
}
type numArray []num

type resultArray []num

func (s numArray) Len() int {
	return len(s)
}
func (s numArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s numArray) Less(i, j int) bool {
	return s[i].n < s[j].n
}
func (s resultArray) Len() int {
	return len(s)
}
func (s resultArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s resultArray) Less(i, j int) bool {
	return s[i].index < s[j].index
}

func main() {
	var t, i, temp int64
	fmt.Scanf("%d\n", &t)
	var n numArray

	for i = 1; i <= t; i++ {
		fmt.Scanf("%d\n", &temp)
		n = append(n, num{temp, i, 0})
	}
	sort.Sort(n)

	var fibA, fibB int64 = 1, 2
	var sum int64 = 0;

	for i = 0; i < t; i++ {
		var currentMax int64 = n[i].n

		for fibB <= currentMax {
			if fibB % 2 == 0 {
				sum += fibB
			}

			var l = fibB
			fibB += fibA
			fibA = l
		}

		n[i].res = sum

	}

	var k resultArray = resultArray(n)
	sort.Sort(k)

	for i = 0; i < int64(len(k)); i++ {
		fmt.Println(k[i].res)
	}
	//fmt.Println(k)

}
