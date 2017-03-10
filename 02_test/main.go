package main

import "fmt"

const tStr = "Sring is here"

const (
	_ = iota
	A = 35 << 1
	B = 1 << (iota * 10)
)
func main() {
	var i = 42

	fmt.Println("1 - ", tStr)
	fmt.Println("2 - ", &i)
	fmt.Println(A)
	fmt.Println(B)

}
