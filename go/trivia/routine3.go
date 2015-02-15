package main

import "fmt"
import "time"

func f() {
  go g()
  for i := 0; i < 2; i++ {
    fmt.Println(1)
  }
}

func g() {
  time.Sleep(time.Duration(1) * time.Second)
  for i := 0; i < 10; i++ {
    fmt.Println(2)
  }
}
func main() {
  go f()
  time.Sleep(time.Duration(3) * time.Second)
}
