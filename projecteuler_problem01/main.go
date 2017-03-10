package main

import (
	"fmt"
	"strconv"
)

func main() {

	var result int64
	var text string
	var resultArr = make([]int64,0)
	var i int64
	//var j,k int64

	fmt.Scanf("%s\n", &text)
	count, _ := strconv.ParseInt(text, 10, 64);

	for i=1; i <= count; i++ {

		fmt.Scanf("%s\n", &text)
		val, _ := strconv.ParseInt(text, 10, 64)
		result = 0;

		var a, b, c int64

		a = (val-1)%3
		b = (val-1)%5
		c = (val-1)%15

		a = val - 1 - a
		b = val - 1 - b
		c = val - 1 - c

		a = a / 3
		b = b / 5
		c = c / 15

		result = 3*a*(a+1)/2 + 5*b*(b+1)/2 - 15*c*(c+1)/2;

		resultArr = append(resultArr,result)

		//result = 0;
		//
		//i3 := int64((val - 1) / 3)
		//i5 := int64((val - 1) / 5)
		//
		//
		//for j=1; j <= i3; j++ {
		//	result += int64(j) * 3
		//}
		//
		//for k=1; k <= i5; k++ {
		//	temp := k * 5
		//	if temp % 3 != 0 {
		//		result += int64(k) * 5
		//	}
		//}
		//
		//resultArr = append(resultArr,result)

	}
	for _, k := range resultArr {
		fmt.Println(k)
	}


}