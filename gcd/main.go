package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	if b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	args := os.Args[1:]

	if len(args) > 2 || len(args) < 2 {
		return
	}

	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])

	//
	res := gcd(a, b)
	fmt.Println(res)
}
