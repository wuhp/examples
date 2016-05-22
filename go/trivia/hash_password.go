package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

func main() {
    fmt.Println(hex.EncodeToString(hmac.New(sha256.New, []byte("0123456789abcdefghijklmnopqrstuv")).Sum([]byte("password"))))
}
