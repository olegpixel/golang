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

func divCount(n int64) int64 {
	var j int64 = 2
	var tt int64
	var result int64 = 1

	for n > 1 {
		tt = 0
		for n % j == 0 {
			tt++
			n /= j
		}
		j++
		if tt > 0 {
			result = result * (tt + 1)
		}
	}

	return result
}

func main() {
	var ts, i, temp, k int64
	var n numArray

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%d\n", &temp)
		n = append(n, num{temp, i, 0})
	}
	sort.Sort(n)

	var premNum int64 = 1;
	var nextInt int64 = 2;
	var dividerCounter int64 = 2


	for i = 0; i < ts; i++ {
		k = n[i].n

		if (k == 0) {
			n[i].res = 1
			continue
		}
		if (k == 1) {
			n[i].res = 3
			continue
		}

		for dividerCounter <= k {

			dividerCounter = 2
			premNum += nextInt
			nextInt++

			dividerCounter = divCount(premNum)

			//fmt.Println("premNum", premNum)
			//fmt.Println("dividerCounter", dividerCounter)
		}

		//fmt.Println(premNum)
		//fmt.Println(dividerCounter)
		n[i].res = premNum

	}

	var km resultArray = resultArray(n)
	sort.Sort(km)

	for i = 0; i < int64(len(km)); i++ {
		fmt.Println(km[i].res)
	}
}
