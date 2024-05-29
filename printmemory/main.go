package main

import (
	"github.com/01-edu/z01"
)

func toHex(n int) string {
	str := ""
	if n == 0 {
		return "00"
	}
	hexvals := "0123456789abcdef"
	for n > 0 {
		str = string(hexvals[n%16]) + str
		n /= 16
	}
	return str
}
func PrintMemory(arr [10]byte) {
	mem := []string{}
	for _, bt := range arr {
		hx := toHex(int(bt))
		mem = append(mem, hx)
	}

	for i := 0; i < len(mem); i++ {
		for _, char := range mem[i] {
			z01.PrintRune(char)
		}
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}

	}
	z01.PrintRune('\n')
	for _, char := range arr {
		if char >= 32 && char <= 126 {
			z01.PrintRune(rune(char))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
