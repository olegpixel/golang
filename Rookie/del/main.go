package main

import (
	"fmt"
	"sort"
	"strings"
)

type word struct {
	str string
	sum uint64
}

type dictionary []word

func (s dictionary) Len() int {
	return len(s)
}
func (s dictionary) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s dictionary) Less(i, j int) bool {
	return s[i].str < s[j].str
}

func calcSum(s string) uint64 {
	var result uint64
	for j := 0; j < len(s); j++ {
		result += uint64(s[j])
	}
	return result

}

func main() {
	var ts, i, result, j uint64
	var temp string
	var input dictionary

	fmt.Scanf("%d\n", &ts)

	for i = 1; i <= ts; i++ {
		//fmt.Scanf("%s", &temp)
		temp = ""
		j = i
		for j != 0 {
			temp += string((j % 25) + 65)
			j = j / 25
		}

		//temp
		input = append(input, word{temp, calcSum(temp)})
	}

	sort.Sort(input)

	for i = 0; i < ts-1; i++ {
		if input[i].sum != 1 {
			if strings.HasPrefix(input[i].str, input[i+1].str) {
				input[i+1].sum = 1;
			}
		}
	}
	for i = 0; i < ts; i++ {
		if input[i].sum != 1 {
			result += input[i].sum
		}
	}
	fmt.Println(input)

	fmt.Println(result)



}
