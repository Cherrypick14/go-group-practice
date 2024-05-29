package main

import "fmt"

func Split(s, sep string) []string {
	arr := []string{}
	sepLen := len(sep)
	pos := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i:i+sepLen] == sep {
			arr = append(arr, s[pos:i])
			pos = i + sepLen
		}
	}
	arr = append(arr, s[pos:len(s)-1])
	return arr
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
