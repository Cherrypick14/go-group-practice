// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
}

func SortWordArr(a []string) {
	slice := []rune{}
	for _, char := range a {
		slice = append(slice, []rune(char)...)
	}

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	fmt.Print(string(slice))
}
