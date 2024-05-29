package main

import (
  "os"
  "fmt"
)

func options(arr []string){
  options:="******zy xwvutsrq ponmlkji hgfedcba"

  mp := map[rune]rune{}

  strConcat:=""
  for _,str := range arr{
    strConcat+=str
  }

  for i,char :=range strConcat{
    if char == '-' && strConcat[i+1] == 'h'{
      fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
      return
    }
    
    if char != '-'{
      _,ok:=mp[char]
      if !ok{
        mp[char] = char
      }

    }
  }
  fmt.Println(mp)
  for i:=0;i<len(options);i++{
    _,ok:=mp[rune(options[i])]

    if ok{
      fmt.Print(1)
    }else{
      if options[i] == ' '{
        fmt.Print(" ")
      }else{
        fmt.Print(0)
      }
    }





  }
  fmt.Println()
}

func main(){

  args := os.Args[1:]

  if len(args) == 0{
    fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
  }

  options(args)
}
