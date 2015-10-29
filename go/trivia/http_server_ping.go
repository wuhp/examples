package main

import (
    "net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
}

func main() {
    http.HandleFunc("/ping", ping)
    http.ListenAndServe(":8080", nil)
}
