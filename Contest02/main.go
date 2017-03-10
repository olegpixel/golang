package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	var i, a, x, num, result, m, counter int64

	fmt.Scanf("%s\n", &input)
	num, _ = strconv.ParseInt(input, 10, 64);

	//num = 100000
	//x = 0

	for i=1; i <= num; i++ {
		fmt.Scanf("%s\n", &input)
		x, _ = strconv.ParseInt(input, 10, 64);

		//x = int64(rand.Intn(1000000000))

		counter = 0
		a = x

		 for (a > 0) {
			a = a >> 1
			 counter++
		 }

		result = 1

		for m=1; m <= counter; m++ {
			result *= 2
		}
		result = result - 1 - x

		fmt.Println(result)

	}

	//fmt.Println(rand.Intn(100))
}
