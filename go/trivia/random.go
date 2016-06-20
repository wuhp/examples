package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Call Seed, using current nanoseconds.
    rand.Seed(int64(time.Now().Nanosecond()))
    // Random int will be different each program execution.
    value := rand.Int()
    fmt.Println(value)
}
