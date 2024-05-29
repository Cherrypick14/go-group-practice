package main

import "fmt"

func SwapBits(octet byte) byte {
	return (octet << 4) | (octet >> 4)
}

func main() {
	var octet byte = 65 // 01000001
	fmt.Printf("Original: %08b\n", octet)
	swapped := SwapBits(octet)
	fmt.Printf("Swapped: %08b\n", swapped)
}
