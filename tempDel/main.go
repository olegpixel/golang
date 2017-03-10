package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i, j, k, temp, result, oneDigit, twoDigits, threeDigits int64


	for ii:=888888888; ii < 888889000; ii++ {

		longStr := strconv.Itoa(ii)
		result = 0
		length := int64(len(longStr))

		temp = 1
		for i = 0; i < length; i++ {

			if (i > 0) {
				temp *= 2
				temp = temp % 1000000007
			}

			//fmt.Println("i = ", i)
			//fmt.Println("temp = ", temp)
			oneDigit, _ = strconv.ParseInt(string(longStr[i]), 10, 64);

			if (oneDigit == 8 || oneDigit == 0) {
				result++
				//fmt.Println("ONE DIGIT  -  ", oneDigit)
			}

			for j = i + 1; j < length; j++ {
				twoDigits, _ = strconv.ParseInt(string(longStr[i]) + string(longStr[j]), 10, 64);

				if (twoDigits % 8 == 0) {
					result++
					//fmt.Println("TWO DIGIT  -  ", twoDigits)
				}

				for k = j + 1; k < length; k++ {
					threeDigits, _ = strconv.ParseInt(string(longStr[i]) + string(longStr[j]) + string(longStr[k]), 10, 64);

					if (threeDigits % 8 == 0) {
						//tempR := result;

						result += temp
						//result++
						//
						//if (i > 0) {
						//	result++
						//}
						//temp = 1
						//for lt = 1; lt < i; lt++ {
						//	temp = temp * 2
						//	temp = temp % 1000000007
						//	result += temp
						//	result = result % 1000000007
						//}

						//fmt.Println("THREE DIGIT  -  ", threeDigits)
						//fmt.Println("result + ", result-tempR)
						//fmt.Println("i = ", i)
					}
					result = result % 1000000007
				}
				result = result % 1000000007
			}
			result = result % 1000000007

		}
		result = result % 1000000007
		fmt.Println("ii = ", longStr)
		fmt.Println("result = ", result)

	}
}
