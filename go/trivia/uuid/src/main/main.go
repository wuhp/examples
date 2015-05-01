package main

import (
    "fmt"
    "github.com/satori/go.uuid"
)

func main() {
    // Creating UUID Version 4
    u1 := uuid.NewV4()
    fmt.Printf("UUIDv4: %s\n", u1.String())
}
