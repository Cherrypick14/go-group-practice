package main

import (
  "os"
  "github.com/01-edu/z01"
)


func itoa(i int)string{
  res:=""
  if i == 0{
    return "0"
  }
  for i>0{
    res = string(i*10 + '0') + res
    i/=10
  }
  return res
}


func brainfuck(s string){



  pointer:=0 
  memory := [2048]byte{}
  index:=0




  for ;index<len(s);index++{
    if s[index] == '>'{
      pointer++
    }

    if s[index] == '+'{
      memory[pointer]++
    }


    if s[index] =='-'{
      memory[pointer]--
    }


    if s[index] == '<'{
      pointer--
    }


    if s[index] == '.'{
      z01.PrintRune(rune(memory[pointer]))
    }

    if s[index] == '['{
      if memory[pointer] == 0{
        loop:=1
        for loop > 0{
          index++
          if s[index] == '[' {
						loop++
					} else if s[index] == ']' {
						loop--
					}
        }
      }
    }

    if s[index] == ']'{
      if memory[pointer] != 0{
        loop:=1
        for loop > 0{
          index--
          if s[index] == '['{
            loop--
          }else if s[index] == ']'{
            loop++
          }
        }
      }
    }
  }
}
func main(){
  args := os.Args[1:]

  if len(args) != 1{
    return
  }

  brainfuck(args[0])
}
