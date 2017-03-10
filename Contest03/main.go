package main

import (
	"fmt"
	"strconv"
)
func main() {
	var input, longStr string
	var i, j, k, temp, result, oneDigit, twoDigits, threeDigits int64

	fmt.Scanf("%s\n", &input)
	//n, _ = strconv.ParseInt(input, 10, 64);

	fmt.Scanf("%s\n", &longStr)

	length := int64(len(longStr))

	temp = 1
	for i = 0; i < length; i++ {

		if (i > 0) {
			temp *= 2
			temp = temp % 1000000007
		}

		//fmt.Println("i = ", i)
		//fmt.Println("temp = ", temp)
		oneDigit, _ =  strconv.ParseInt(string(longStr[i]), 10, 64);

		if (oneDigit == 8 || oneDigit == 0) {
			result++
			//fmt.Println("ONE DIGIT  -  ", oneDigit)
		}

		for j = i+1; j < length; j++ {
			twoDigits, _ =  strconv.ParseInt(string(longStr[i])+string(longStr[j]), 10, 64);

			if (twoDigits % 8 == 0) {
				result++
				//fmt.Println("TWO DIGIT  -  ", twoDigits)
			}

			for k = j+1; k < length; k++ {
				threeDigits, _ =  strconv.ParseInt(string(longStr[i])+string(longStr[j])+string(longStr[k]), 10, 64);

				if (threeDigits % 8 == 0) {
					tempR := result;
					result += temp
					fmt.Println("THREE DIGIT  -  ", threeDigits)
					fmt.Println("result + ", result-tempR)
				}
				result = result % 1000000007
			}
			result = result % 1000000007
		}
		result = result % 1000000007

	}
	result = result % 1000000007
	fmt.Println(result)
}
