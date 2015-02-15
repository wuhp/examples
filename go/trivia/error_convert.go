package main

import "errors"
import "fmt"

func main() {
  err := errors.New("aaa")
  fmt.Println(err)
  a := []byte(err.Error())
  fmt.Println(a)
}
