package main

import "fmt"

func main() {
  a := [...]int{1,2,3}
  s1 := a[1:2]
  s2 := append(s1, 10)
  a[1] = 100
  fmt.Println(cap(s1))
  fmt.Println(a)
  fmt.Println(s1)
  fmt.Println(s2)
}
