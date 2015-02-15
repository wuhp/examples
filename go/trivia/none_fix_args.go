package main

import "fmt"

func main() {
  f1(1, 2, 3, 4)
}

func f1(arg ...int) {
  fmt.Println(arg)
  f2(arg...)
  f3(arg[0:1]...)
}

func f2(arg ...int) {
  return
}

func f3(arg ...int) {
  return
}
