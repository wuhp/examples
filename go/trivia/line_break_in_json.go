package main

import (
    "encoding/json"
    "fmt"
)

type A struct {
    S string `json:"s"`
}

func main() {
/*
    This is not supported, use \n to replace a real line break
    s := `{"s":"a
c"}`
*/
    s := `{"s":"a\nc"}`
    a := new(A)
    if err := json.Unmarshal([]byte(s), a); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(a.S)
}
