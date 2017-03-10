package main

import "fmt"

const div int = 600851475143

func main() {

	temp := div

	for i:=3; (temp > 1); i = i + 2 {
		//fmt.Println(i)
		for temp % i == 0 {
			fmt.Println(i)
			temp = temp / i
		}
	}
	fmt.Println(temp)
}
