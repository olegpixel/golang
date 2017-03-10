package main

import (
	"fmt"
)

//type arrayAndInt struct {
//	number int64
//	result bool
//}

var possibleNumbers = make(map[int64]bool)

//var answer = make([]arrayAndInt, 1, 100)

func lookFor(n int64) {
	if (!possibleNumbers[n]) {
		possibleNumbers[n] = true

		//for i:=1; i < len(answer); i++ {
		//	if answer[i].number == n {
		//		answer[i].result = true
		//	}
		//}
	}
}

func main() {
	var count, i, input int64
	var str string
	//var max int64
	//var min int64 = math.MaxInt64

	fmt.Scanf("%v\n", &str)
	fmt.Scanf("%d\n", &count)

	var sum int64 = int64(str[0]) - 96
	//if sum >= min && sum <= max  {
		lookFor(sum)
	//}

	for i = 1; i < int64(len(str)); i++ {
		if str[i] == str[i-1] {
			sum += int64(str[i]) - 96
		} else {
			sum = int64(str[i]) - 96
		}

		//if sum >= min && sum <= max  {
			lookFor(sum)
		//}
	}

	//fmt.Println(possibleNumbers)
	for i = 1; i <= count; i++ {
		fmt.Scanf("%d\n", &input)

		if possibleNumbers[input] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		//temp := arrayAndInt{input, false}
		//if input > max {
		//	max = input
		//}
		//if input < min {
		//	min = input
		//}
		//answer = append(answer, temp)
	}

	//for i:=1; i < len(answer); i++ {
	//	if answer[i].result {
	//		fmt.Println("Yes")
	//	} else {
	//		fmt.Println("No")
	//	}
	//}
}
