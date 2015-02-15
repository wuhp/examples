package main

import "fmt"
import "encoding/json"
import "io/ioutil"

type A struct {
  X int    `json:"aaa"`
  Y string `json:"bbb"`
  Z []int  `json:"ccc"`
}

func main() {
  b, e := ioutil.ReadFile("a.json")
  if e != nil {
    fmt.Println(e)
  }
  fmt.Println(string(b))

  var a A
  e = json.Unmarshal(b, &a)
  if e != nil {
    fmt.Println(e)
  }

  fmt.Println(a.X)
  fmt.Println(a.Y)
  fmt.Println(a.Z)
}
