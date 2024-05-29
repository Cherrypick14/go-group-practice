package main

import (
	//"fmt"
	"os"
	//"strconv"
  "github.com/01-edu/z01"
)

func atoi(s string)int{
  res:=0
  for _,c :=range s{
    res = res*10 + int(c-'0')
  }
  return res
}

func itoa(n int) string{
  res:=""
  if n ==0 {
    return "0"
  }

  for n >0{
    res =string(n%10 + '0') +res
    n/=10
  }
  return res
}



func TabMult(s string) {
  for i:=1;i<=9;i++{
    str:= itoa(i)+" "+"x"+" "+s+" "+"="+" " + itoa(i*atoi(s))

    for _,char :=range str{
      z01.PrintRune(char)
    }
    z01.PrintRune('\n')
  }
}

func main() {
	args := os.Args[1:]
	if len(args)-1 > 0 || len(args)-1 < 0 {
		return
	}
	TabMult(args[0])
}
