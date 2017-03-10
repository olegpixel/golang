package main

import "fmt"

var x int

func increment() {
	x++
	//return x
}
func main() {
	increment()
	fmt.Printf("%d",x)
}
