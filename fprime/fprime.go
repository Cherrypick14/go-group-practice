package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findNextPrime(n int) int {
	for !isPrime(n) {
		n++
	}
	return n
}

func factors(n int) []int {
	if n == 1 {
		return nil
	}
	nextPrime := findNextPrime(2)
	for n%nextPrime != 0 {
		nextPrime = findNextPrime(nextPrime + 1)
	}
	return append([]int{nextPrime}, factors(n/nextPrime)...)
}

func main() {
	x := factors(10000)
	fmt.Println(x)
}
