package main

import "strconv"
import "fmt"

func main() {
  var id int64
  id, _ = strconv.ParseInt("123", 10, 64)
  fmt.Println(id)
}
