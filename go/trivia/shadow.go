package main

import "fmt"

func f() (int, int) {
  return 1,2
}

func g() (a int, b int) {
  x, b := f()
  a = x
  if a == 3 {
    return
  }
  fmt.Println(a)
  fmt.Println(b)
}

func main() {
  x, y := g()
  fmt.Println(x)
  fmt.Println(y)
}
