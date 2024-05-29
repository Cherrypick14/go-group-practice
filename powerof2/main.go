package main

import "fmt"

func isPower(n int) bool {
	for i := 1; i <= n; i *= 2 {
		if i == n {
			return true
		}
	}
	return false
}

func main() {
	x := isPower(4)
	fmt.Println(x)
}
