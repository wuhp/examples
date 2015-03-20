package main

import (
  "net"
  "fmt"
  "os"
)

func handleConnection(conn net.Conn, i int) {
  defer conn.Close()

  fmt.Fprintf(os.Stdout, "Connect succeed! ID: %d\n", i)
  res := fmt.Sprintf("You are the %dth client", i)
  n, err := conn.Write([]byte(res))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Write error: %s\n", err.Error())
  } else {
    fmt.Fprintf(os.Stdout, "Write back %d bytes to client %d\n", n, i)
  }
}

func main() {
  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
    fmt.Fprintf(os.Stderr, "Listen error: %s\n", err.Error())
    os.Exit(1)
  }

  i := 1
  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Fprintf(os.Stderr, "Accept error: %s\n", err.Error())
      continue
    }

    i += 1
    go handleConnection(conn, i)
  }
}
