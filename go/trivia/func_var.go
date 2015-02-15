package main

import "fmt"

func f(a int) func(int) int {
  x := func(y int) int {
    return a + y
  }
  return x
}

func main() {
  m := f(1)
  fmt.Printf("%d, %d\n", m(2), m(100))
}
