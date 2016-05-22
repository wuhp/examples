package main

import (
    "fmt"
)

type A struct {
    X int
}

func main() {
    a1 := A{X: 1}
    a2 := a1
    fmt.Println(a1)
    fmt.Println(a2)
}
