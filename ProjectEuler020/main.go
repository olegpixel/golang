package main

import (
	"fmt"
	"sort"
)


type num struct {
	n uint64
	index uint64
	res uint64
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

func digitsDum(a []uint64) uint64 {
	var s, u uint64
	tmp := make([]uint64, len(a))
	copy(tmp, a)

	for u = 0; u < uint64(len(tmp)); u++ {
		for tmp[u] != 0 {
			s += tmp[u] % 10
			tmp[u] /= 10
		}
	}

	return s
}


func main() {
	var ts, counter, i, j, temp, k uint64
	var n numArray
	var numberA []uint64

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		fmt.Scanf("%d\n", &temp)
		n = append(n, num{temp, i, 0})
	}
	sort.Sort(n)

	counter = 0
	numberA = append(numberA, 1)

	for i = 0; i <= ts-1; i++ {
		k = n[i].n

		//fmt.Println("num before = ",numberA)
		for j = counter+1; j <= k; j++ {
			var ij = j
			for ij % 10 == 0 && ij != 0{
				ij /= 10
			}
			//fmt.Println("j=",j)

			temp2 := len(numberA)
			for temp = 0; temp < uint64(temp2); temp++ {
				//if number[temp] != 0 {
					numberA[temp] *= ij
				//} else {
				//	number[temp] = ij
				//}
			}
			//fmt.Println("num temp = ", numberA)
			for temp = 0; temp < uint64(len(numberA)); temp++ {
				if numberA[temp] >= 10 {
					var z = numberA[temp] / 10
					numberA[temp] = numberA[temp] % 10
					if temp < uint64(len(numberA))-1 {
						numberA[temp+1] += z
					} else {
						numberA = append(numberA, z)
					}
					temp--
				}
			}
		}
		//fmt.Println("num after = ",numberA)
		n[i].res = digitsDum(numberA)
		counter = k
	}

	var km resultArray = resultArray(n)
	sort.Sort(km)

	for i = 0; i < uint64(len(km)); i++ {
		fmt.Println(km[i].res)
	}
}
