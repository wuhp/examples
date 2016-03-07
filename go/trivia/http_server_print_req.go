package main

import (
    "fmt"
    "os"
    "io"
    "bytes"
    "net/http"
)

var n int64 = 0

func parse(w http.ResponseWriter, r *http.Request) {
    n++

    fmt.Printf("\n--- request %d ---\n", n)
    fmt.Printf("RemoteAddr: %s\n", r.RemoteAddr)
    fmt.Printf("Host: %s\n", r.Host)
    fmt.Printf("ContentLength: %d\n", r.ContentLength)
    fmt.Printf("Proto: %s\n", r.Proto)
    fmt.Printf("Method: %s\n", r.Method)
    fmt.Printf("URL: %#v\n", r.URL)

    fmt.Printf("Header:\n")
    for k, _ := range r.Header {
        fmt.Printf("  %s: %s\n", k, r.Header.Get(k))
    }

    defer r.Body.Close()
    var b bytes.Buffer
    io.Copy(&b, r.Body)
    fmt.Printf("Body: %s\n", string(b.Bytes()))
}

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("Usage: %s <port>\n", os.Args[0])
        return
    }

    fmt.Printf("Listening on 0.0.0.0:%s\n", os.Args[1])

    http.HandleFunc("/", parse)
    http.ListenAndServe(fmt.Sprintf(":%s", os.Args[1]), nil)
}
