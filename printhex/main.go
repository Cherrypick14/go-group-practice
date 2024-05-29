package main

import (
	"fmt"
	"os"
	"strconv"
)

func printhex(s string) {
	hexVals :="0123456789abcdef"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error")
		return
	}
	hexString := ""

	for i > 0 {
		hexString +=  string(hexVals[i%16])
		i = i / 16
	}
	fmt.Println(hexString)
}

func main() {
	args := os.Args[1:]

	s := args[0]
	printhex(s)
}
