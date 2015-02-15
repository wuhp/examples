package main

import "fmt"

func main() {
  fmt.Println(1)
  c := make(chan bool)
  go func() {
    fmt.Println(101)
    //close(c)
    c <- true
    fmt.Println(102)
  }()
  fmt.Println(2)
  <- c
  fmt.Println(3)
}
