package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
)

func main() {

	var text1 string

	fmt.Scanf("%s", &text1)

	text := strings.TrimSpace(text1)
	i, _ := strconv.ParseInt(text, 10, 64);

	var a,b,c int64
	var positive, negative, zeros float64

	arr := bufio.NewReader(os.Stdin)
	line, err := arr.ReadString('\n')
	line = strings.TrimSpace(line)
	//fmt.Println(line)
	//fmt.Println(err)
	for strings.Contains(line, "  ") {
		line = strings.Replace(line, "  ", " ", -1)
	}
	strs := strings.Split(line, " ")

	//fmt.Println(strs)
	nums := make([]float64, len(strs))
	for i, str := range strs {
		if nums[i], err = strconv.ParseFloat(str, 64); err != nil {
			log.Fatal(err)
		}
	}

	for j := 0; j < len(nums); j++ {
		if nums[j] > 0 {
			a++
		} else if nums[j] < 0 {
			b++
		} else {
			c++
		}
	}
	positive = float64(a) / float64(i)
	negative = float64(b) / float64(i)
	zeros =  float64(c) / float64(i)
	fmt.Println(positive)
	fmt.Println(negative)
	fmt.Println(zeros)
}
