package main

import (
	"fmt"
	"os"
)

func lastWord(s string) string {
	str := ""
	arr := []string{}

	for _, char := range s {
		if char != ' ' {
			str += string(char)
		}
		if char == ' ' && str != "" {
			arr = append(arr, str)
			str = ""
		}
	}
	if str != "" {
		arr = append(arr, str)
	}
	return arr[len(arr)-1]
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	fmt.Println(lastWord(args[0]))
}
