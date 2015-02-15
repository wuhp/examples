package main

import "fmt"

func f() {
  fmt.Println(1)
}


func main() {
  a := f
  a()
}
