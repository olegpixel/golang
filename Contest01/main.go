package main

import (
	"fmt"
	"strconv"
)

func main() {

	result := true

	var inputStr string
	fmt.Scanf("%s", &inputStr)
	num, _ := strconv.ParseInt(inputStr, 10, 64);

	fmt.Scanf("%s", &inputStr)
	m, _ := strconv.ParseInt(inputStr, 10, 64);

	fmt.Scanf("%s", &inputStr)
	c, _ := strconv.ParseInt(inputStr, 10, 64);

	volume := c * m

	var i int64
	for i=1; i <= num; i++ {
		fmt.Scanf("%s", &inputStr)
		vl, _ := strconv.ParseInt(inputStr, 10, 64);

		if (vl > volume) {
			result = false;
		}
	}
	if (result) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
