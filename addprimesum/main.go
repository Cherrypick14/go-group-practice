package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nextPrimesum(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) == true {
			sum += i
		}
	}
	return sum
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 1 {
		fmt.Println("0")
		return
	}
	x, _ := strconv.Atoi(args[0])
	if x < 1 {
		fmt.Println("0")
		return
	}

	a := nextPrimesum(x)
	fmt.Println(a)
}
