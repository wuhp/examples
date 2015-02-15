package main

import "fmt"

func main() {
  a := []int{1, 2, 3, 4}
  b := a[:]
  b = append(b[:1], b[2:]...)
  fmt.Println(b)
}
