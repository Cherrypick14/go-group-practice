package main

import (
	"errors"
	"fmt"
	"os"
)

type Stack struct {
	content []int
}

func push(stack *Stack, n int) {
	stack.content = append(stack.content, n)
}

func pop(stack *Stack) int {
	val := stack.content[len(stack.content)-1]
	stack.content = append([]int{}, stack.content[:len(stack.content)-1]...)
	return val
}

func toArr(s string) []string {
	arr := []string{}
	str := ""
	for _, char := range s {
		if char != ' ' {
			str += string(char)
		} else if str != "" && char == ' ' {
			arr = append(arr, str)
			str = ""
		}
	}
	if str != "" {
		arr = append(arr, str)
	}
	return arr
}
func atoi(s string) (int, error) {
	res := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			res = res*10 + int(char-'0')
		} else {
			return 0, errors.New("Error")
		}
	}
	return res, nil
}
func reversePolishNotation(s string) (int, error) {
	str := toArr(s)
	stack := &Stack{}
	for _, val := range str {
		if val != "*" && val != "-" && val != "+" && val != "/" && val != "%" {
			x, err := atoi(val)
			if err != nil {
				return 0, errors.New("Error")
			}
			push(stack, x)
		} else if val == "+" {
			x := pop(stack)
			y := pop(stack)
			calc := y + x
			push(stack, calc)
		} else if val == "-" {
			x := pop(stack)
			y := pop(stack)
			calc := y - x
			push(stack, calc)

		} else if val == "*" {
			x := pop(stack)
			y := pop(stack)
			calc := y * x
			push(stack, calc)

		} else if val == "%" {
			x := pop(stack)
			y := pop(stack)
			calc := y % x
			push(stack, calc)
		} else if val == "/" {
			x := pop(stack)
			y := pop(stack)
			calc := y / x
			push(stack, calc)
		}
	}
	return pop(stack), nil
}
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	x, err := reversePolishNotation(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(x)
}
