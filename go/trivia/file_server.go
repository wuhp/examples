package main

import (
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("/tmp/test_file_server")))
    http.ListenAndServe(":8108", nil)
}
