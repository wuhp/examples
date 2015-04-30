package main

import (
  "time"
  "fmt"
  "log"
  "os"
  "strings"

  "github.com/apcera/nats"
)

func main() {
  args := os.Args
  if len(args) != 3 {
    fmt.Println("my-pub <topic> <msg>")
    return
  }

  opts := nats.DefaultOptions
  opts.Servers = []string{"nats://localhost:4222"}
  for i, s := range opts.Servers {
    opts.Servers[i] = strings.Trim(s, " ")
  }

  nc, err := opts.Connect()
  if err != nil {
    log.Fatalf("Can't connect nats-server: %v\n", err)
  }
  defer nc.Close()

  log.Printf("Requested [%s] : '%s'\n", args[1], args[2])
  msg, err := nc.Request(args[1], []byte(args[2]), 1 * time.Second)
  if err != nil {
    if err == nats.ErrTimeout {
      log.Printf("Timeout\n")
    } else {
      log.Printf("Internal error")
    }
    return
  }

  log.Printf("Received [%s]\n", string(msg.Data))
}
