package main

import "fmt"

func Rot14(s string) string {
	str := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			str += string((char-'A'+14)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			str += string((char-'a'+14)%26 + 'a')
		} else {
			str += string(char)
		}
	}
	return str
}

func main() {
	fmt.Println(Rot14("Hello! How are You?"))
}
