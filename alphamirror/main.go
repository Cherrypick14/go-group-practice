package main

import "fmt"

// func AlphaMirror(s string) string {
// 	str := ""

// 	for _, char := range s {
// 		if char >= 'A' && char <= 'Z' {
// 			pos := (char + 25) % 90
// 			str += string(90 - pos)
// 		} else if char >= 'a' && char <= 'z' {
// 			pos := (char + 25) % 122
// 			str += string(rune(122 - pos))
// 		} else {
// 			str += string(char)
// 		}
// 	}
// 	return str
// }

func AlphaMirror(s string) string {
	str := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			str += string((90-char)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			str += string((122-char)%26 + 'a')
		} else {
			str += string(char)
		}
	}
	return str
}

func main() {
	x := AlphaMirror("My horse is Amazing.")
	fmt.Println(x)
}
