package main

import (
	"fmt"
	"os"
)

func hiddenp(a, b string) string {
	s := ""
	pointer := 0
	for i := 0; i < len(b); i++ {
		if pointer < len(a) {
			if a[pointer] == b[i] {
				s += string(b[i])
				pointer++
			}
		}
	}

	if s == a {
		return "1"
	}
	return "0"
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 || len(args) > 2 {
		return
	}
	res := hiddenp(args[0], args[1])
	fmt.Println(res)
}
