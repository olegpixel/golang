package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func main() {
	var input, longStr string
	var i, j, k, temp, lt, result, length uint32
	var oneDigit, twoDigits, threeDigits uint16

	fmt.Scanf("%s\n", &input)
	//n, _ = strconv.ParseInt(input, 10, 64);

	fmt.Scanf("%s\n", &longStr)

	length = uint32(len(longStr))

	longStr = reverse(longStr)

	//temp = 1
	for i = 0; i < length; i++ {

		oneDigitT, _ :=  strconv.ParseUint(string(longStr[i]), 10, 32);

		oneDigit = uint16(oneDigitT)

		if (oneDigit == 8 || oneDigit == 0) {
			result++
		}
		if (oneDigit % 2 != 0) {
			continue
		}

		for j = i+1; j < length; j++ {
			twoDigitsT, _ :=  strconv.ParseUint(string(longStr[j]) + string(longStr[i]), 10, 32);

			twoDigits = uint16(twoDigitsT)

			if (twoDigits % 8 == 0) {
				result++
			}

			if (twoDigits % 4 != 0) {
				continue
			}

			for k = j+1; k < length; k++ {
				threeDigitsT, _ :=  strconv.ParseUint(string(longStr[k])+string(longStr[j])+string(longStr[i]), 10, 32);

				threeDigits = uint16(threeDigitsT)

				if (threeDigits % 8 == 0) {

					result++

					temp = 1
					for lt = 1; lt <= length-k-1; lt++ {
						result += temp
						result = result % 1000000007
						temp = temp * 2
						temp = temp % 1000000007
					}

				}
				//result = result % 1000000007
			}
			result = result % 1000000007
		}
		result = result % 1000000007

	}
	//result = result % 1000000007
	fmt.Println(result)
}
