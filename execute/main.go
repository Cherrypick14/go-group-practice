package main

import (
	"fmt"
	"os/exec"
)

func main() {
	x := exec.Command("go", "run", "./hello/hello.go")
	b, _ := x.Output()
	fmt.Print(string(b))
}
