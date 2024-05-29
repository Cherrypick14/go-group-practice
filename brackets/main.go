package main

import (
	"fmt"
	"os"
)

func isCorrectlyBracketed(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 || brackets[r] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	for _, arg := range os.Args[1:] {
		if isCorrectlyBracketed(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}

