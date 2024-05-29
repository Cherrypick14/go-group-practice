package main

import "fmt"

func fib(i int, m map[int]int) int {
	if i < 0 {
		return -1
	}
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	val, ok := m[i]
	if ok {
		return val
	}
	m[i] = fib(i-1, m) + fib(i-2, m)

	return fib(i-1, m) + fib(i-2, m)
}

func main() {
	a := 900
	x := map[int]int{}
	b := fib(a, x)
	fmt.Println(b)
}
