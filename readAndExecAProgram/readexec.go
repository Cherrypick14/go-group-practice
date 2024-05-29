package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	file, _ := os.ReadFile("./readAndExecAProgram/main.txt")
	outputfile, err := os.Create("./readAndExecAProgram/output.go")
	if err != nil {
		fmt.Println(err)
	}
	outfile, err := os.Create("./readAndExecAProgram/output.txt")
	if err != nil {
		fmt.Println(err)
	}
	//

	outputfile.WriteString(string(file))

	//
	a := exec.Command("go", "run", "./readAndExecAProgram/output.go")
	x, err := a.Output()
	if err != nil {
		fmt.Println(err)
		outfile.WriteString(string(err.Error()))
	}
	outfile.WriteString(string(x))
	outfile.Close()
	os.Remove("./readAndExecAProgram/output.go")
}
