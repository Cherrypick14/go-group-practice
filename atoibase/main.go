package main

import "fmt"

func AtoiBase(s string, t string) int {
	res := 0
	lens := len(t)
	if lens <= 1 {
		return res
	}

	arr := make(map[rune]int)
	for i, c := range t {

		_, exists := arr[c]
		if exists || c == '-' || c == '+' {
			return res
		}

		// if _, exists := arr[c]; exists || c == '-' || c == '+' {
		// 	return res
		// }
		arr[c] = i
	}

	for _, v := range s {
		if val, exists := arr[v]; exists {
			res = res*lens + val
		}
	}
	fmt.Println(arr)
	return res
}

// }

func main() {
	// x := AtoiBase("7d", "0123456789abcdef")
	// fmt.Println(x)
	// x = AtoiBase("uoi", "choumi")
	// fmt.Println(x)
	// fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("1255", "0123456789"))
}
