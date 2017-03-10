package main

import "fmt"

func main() {
	var count, input, i int64

	fmt.Scanf("%d\n", &count)

	for i=1; i <= count; i++ {
		fmt.Scanf("%d\n", &input)

		if (input <= 37) {
			fmt.Println(input)
		} else if (input % 5 == 3 || input % 5 == 4) {
			fmt.Println((input / 5 + 1) * 5)
		} else {
			fmt.Println(input)
		}
	}
}
