package main

import (
	"fmt"
	"strconv"
)

var numbers2Dig[100] uint64
var numbers1Dig[10] uint64
var powers[][] uint64

var result uint64

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func addNum2(n int8) {
	var i int8
	for i=0; i<10; i += 2 {
		if (numbers1Dig[i] > 0) {
			numbers2Dig[n * 10 + i] += numbers1Dig[i]

			if ((n * 10 + i) % 8 == 0) {
				result += numbers1Dig[i]
			}
		}
	}
}

func addNum3(n int8, m uint64) {
	var i int8
	for i=0; i<100; i += 4 {
		if (numbers2Dig[i] > 0) {
			//var temp uint64 = 1
			//var temp2 uint64 = 1
			var temp3 uint64 = uint64(n) * 100 + uint64(i)

			if ( (temp3) % 8 == 0) {
				//result++

				//temp2 = 1
				//var tempR uint64 = 1
				//
				////fmt.Printf("Three dig result + %v, num = %v, n = %v, i = %v\n", temp2,temp3, n, i)
				////fmt.Printf("Result = %v\n", result)
				//
				//
				//for temp = 1; temp < m; temp++ {
				//	tempR += temp2
				//	//fmt.Printf("Three dig result + %v, num = %v, n = %v, i = %v\n", temp2,temp3, n, i)
				//	//fmt.Printf("Result = %v\n", result)
				//	temp2 *= 2
				//	temp2 = temp2 % 1000000007
				//	tempR = tempR % 1000000007
				//}
				//
				//result += uint64(tempR * uint64(numbers2Dig[i])) % 1000000007
				//fmt.Println("result + ", result-tempR)
				powers = append(powers, []uint64{m, numbers2Dig[i]})
			}
		}
	}
}

func main() {

	var input string


	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%s\n", &input)

	input = reverse(input)

	length := uint64(len(input))
	var i uint64

	for i = 0; i < length; i++ {
		i64, _ := strconv.ParseInt(string(input[i]), 10, 64)
		num := int8(i64)

		if (num % 8 == 0) {
			result++
		}

		// append three digits array
		addNum3(num, length - i)

		// append two digits array
		addNum2(num)

		// append one digit array
		if (num % 2 == 0) {
			numbers1Dig[num]++
		}
		result = result % 1000000007
	}

	//fmt.Println(powers)
	var tempR, tempI uint64 = 1, 1

	for len(powers)>0 && tempI < 100000 {

		for i=0; i < uint64(len(powers)); i++ {
			if (powers[i][0] == tempI) {
				result += (tempR * powers[i][1]) % 1000000007
				powers = append(powers[:i], powers[i+1:]...)
				i--
			}
		}
		tempR = tempR % 1000000007
		tempR *= 2
		result = result % 1000000007
		tempI++

	}

	fmt.Println(result)
}
