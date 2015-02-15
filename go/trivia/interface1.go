package main

import "fmt"

type A struct {
  i int
}

func (a A) f1() int {
  return a.i
}

func (a *A) f1() int {
  return a.i
}

func main() {
  var a A
  fmt.Println(a.f1())
//  b := new(A)
//  fmt.Println(a.f1())
}
