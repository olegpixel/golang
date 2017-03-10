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
	var result int64

	fmt.Scanf("%s\n", &text1)

	text := strings.TrimSpace(text1)
	_, err := strconv.ParseInt(text, 10, 64);

	arr := bufio.NewReader(os.Stdin)
	line, err := arr.ReadString('\n')
	line = strings.TrimSpace(line)

	for strings.Contains(line, "  ") {
		line = strings.Replace(line, "  ", " ", -1)
	}
	strs := strings.Split(line, " ")

	nums := make([]int64, len(strs))
	//i = nums[0]
	for i, str := range strs {
		if nums[i], err = strconv.ParseInt(str, 10, 64); err != nil {
			log.Fatal(err)
		}
	}

	for j := 0; j < len(nums); j++ {
		for k:=j+1; k < len(nums); k++ {
			if (nums[j] == nums[k]) && (nums[j] != -1) {
				result++
				nums[j] = -1
				nums[k] = -1
				//fmt.Printf("j = %v\n",j)
				//fmt.Printf("k = %v\n",k)
				break
			}
		}

	}
	fmt.Println(result)
}
