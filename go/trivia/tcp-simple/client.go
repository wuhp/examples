package main

import (
  "fmt"
  "os"
  "net"
  "time"
)

func main() {
  for i := 0; i < 3; i++ {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
      fmt.Fprintf(os.Stderr, "Dial error: %s\n", err.Error())
      os.Exit(1)
    }

    fmt.Fprintf(os.Stdout, "Connect succeed!\n")
    buf := make([]byte, 512)
    n, err := conn.Read(buf)
    fmt.Println(n)
    fmt.Println(err)
    fmt.Fprintf(os.Stdout, "Received response: %s\n", string(buf))
    time.Sleep(1 * time.Second)
    conn.Close()
  }
}
