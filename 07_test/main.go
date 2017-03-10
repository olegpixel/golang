//package main
//
//import "fmt"
//
//func main() {
//	var input string
//
//	fmt.Scanf("%s", &input)
//
//	for shortenString(&input) {
//	}
//	if len(input) == 0 {
//		input = "Empty String"
//	}
//	fmt.Println(input)
//
//}
//
//func shortenString(str *string) bool {
//	var result = false
//
//	for i := 1; i < len(*str); i++ {
//		if (*str)[i] == (*str)[i-1] {
//			if i < len(*str) {
//				*str = (*str)[:i-1] + (*str)[i+1:]
//			} else {
//				*str = (*str)[:i-1]
//			}
//			i--
//			result = true
//		}
//	}
//
//	return result
//
//}


package main
import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for true {
		arr := make([]byte, 0)
		for i := 0; i < len(s); i++ {
			if i + 1 >= len(s) || s[i] != s[i + 1] {
				arr = append(arr, s[i])
			} else {
				i++
			}
			fmt.Printf("i = %v; %s\n",i, arr)
		}

		news := string(arr)

		fmt.Printf("New s = %v\n", news)
		if len(arr) == len(s) {
			s = news
			break
		}

		s = news
	}

	if len(s) == 0 {
		fmt.Println("Empty String")
	} else {
		fmt.Println(s)
	}
}