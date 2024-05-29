package main

import (
	"fmt"
	"os"
)

func union(a, b string) string {
	s1 := map[string]string{}

	a1 := ""
	for _, char := range a + b {
		_, ok := s1[string(char)]
		if !ok {
			s1[string(char)] = string(char)
			a1 += string(char)
		}

	}

	return a1
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 || len(args) > 2 {
		return
	}
	x := union(args[0], args[1])
	fmt.Println(x)
}
