package main

import "fmt"

//func doubleFunc(i int) (float64, bool) {
// return float64(i)/2, i % 2 == 0
//}

func main() {

	var j int
	j = 7

	fn := func(i int) (float64, bool) {
		return float64(i)/2, i % 2 == 0
	}
//	fmt.Scan(j)

	fmt.Println(fn(j));

}
