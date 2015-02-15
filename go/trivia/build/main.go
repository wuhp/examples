package main

import "fmt"
import "math"

func IsPrim(n int) bool {
  for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}

func main() {
  fmt.Println(IsPrim(10))
}
