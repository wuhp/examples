package main

import (
  "fmt"
  "regexp"
)

func main() {
  ok, _ := regexp.MatchString(`\d+`, "123")
  fmt.Println(ok)
}
