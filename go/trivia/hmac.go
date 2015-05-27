package main

import (
    "crypto/hmac"
    "crypto/sha1"
    "crypto/sha256"
    "fmt"
    "io"
)

func main() {
    //sha1
    h := sha1.New()
    io.WriteString(h, "aaaaaa")
    fmt.Printf("%x\n", h.Sum(nil))

    //hmac ,use sha256
    key := []byte("123456")
    mac := hmac.New(sha256.New, key)
    mac.Write([]byte("aaaaaa"))
    fmt.Printf("%x\n", mac.Sum(nil))
}
