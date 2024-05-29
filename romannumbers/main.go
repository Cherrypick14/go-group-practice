package main

import (
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	key   int
	value string
}

var roman = []Item{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// var romanString string = ""

func RomanNumbers(s string) {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	ret := ""
	// operation := []string{}
	ret2 := ""
	for n > 0 {
		for _, v := range roman {
			if n-v.key >= 0 {
				n -= v.key
				ret += v.value

				if len(v.value) > 1 {
					f, s := v.value[1], v.value[0]
					ret2 += "(" + string(f) + "-" + string(s) + ")"
				}
				break
			}
		}
	}

	fmt.Println(ret)
	fmt.Println(ret2)
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}
	RomanNumbers(args[0])
}
