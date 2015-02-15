package main

import "fmt"

func main() {
I:  for i := 0; i < 3; i++ {
    fmt.Printf("loop i = %d\n", i)
    for j := 0; j < 3; j++ {
      if j == 1 {
        continue I
      }
      fmt.Printf("loop j = %d\n", j)
    } 
  }
}
