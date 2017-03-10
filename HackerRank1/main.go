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

	var inputStr string
	var inputInt uint64
	var zeroArr []int
	var counter, result int

	fmt.Scanf("%s", &inputStr)

	inputInt, _ = strconv.ParseUint(inputStr, 10, 0)
	inputBinInt := strconv.FormatUint(inputInt, 2)
	inputBinInt = reverse(inputBinInt)

	fmt.Printf("Binary - %v\n",inputBinInt)

	for i:=0; i < len(inputBinInt); i++ {
		if string(inputBinInt[i]) == "0" {
			counter++
		} else {
			if counter != 0 {
				zeroArr = append(zeroArr, counter)
				//zeroArr = append(zeroArr, 0)
			} else {
				zeroArr = append(zeroArr, 0)
			}
			counter = 0
		}
	}
	//if counter != 0 {
	//	zeroArr = append(zeroArr, counter)
	//}
	fmt.Printf("Array - %v\n", zeroArr)

	result = 1
	for i:=0; i < len(zeroArr); i++ {
		result = result * (zeroArr[i] + 1)
	}

	fmt.Println(result)

	var subResult int
	subResult = 0
	counter = 0
	for i:=0; i < len(zeroArr); i++ {
		if zeroArr[i] != 0 {
			if subResult == 0 {
				subResult = 1
			}
			if counter != 0 {
				subResult *= counter
			}
			subResult *= zeroArr[i]
			counter = 0
		} else {
			counter++
		}
	}
	if counter != 0 {
		//subResult *= counter
	}
	result += subResult
	fmt.Println(result)
}
