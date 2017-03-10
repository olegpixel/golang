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

func primeCheck(a int64) bool {
	var i int64

	if(a == 1) {
		return true;
	}

	for i = 2; i*i <= a; i++ {
		if (a % i == 0) {
			return false;
		}
	}
	return true;
}

func main() {
	var ts, result, i, j, temp, k, prev int64
	var n numArray

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%d\n", &temp)
		n = append(n, num{temp, i, 0})
	}
	sort.Sort(n)


	for i = 0; i <= ts-1; i++ {
		k = n[i].n
		if i == 0 {
			result = 2
			prev = 2
		}

		for j = prev; j <= k; {
			result++
			for !primeCheck(result) {
				result++
			}
			j++
		}

		n[i].res = result
		result--
		prev = k
	}

	var km resultArray = resultArray(n)
	sort.Sort(km)

	for i = 0; i < int64(len(km)); i++ {
		fmt.Println(km[i].res)
	}
}
