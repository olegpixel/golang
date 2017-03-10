package main

import "fmt"

func main() {
	var x [256]string

	for i:=65; i<=255; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
}
