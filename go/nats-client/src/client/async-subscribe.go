package main

import (
  "fmt"
  "log"
  "os"
  "strings"
  "runtime"

  "github.com/apcera/nats"
)

func main() {
  args := os.Args
  if len(args) != 2 {
    fmt.Println(args[0], "<topic1,topic2,...>")
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

  ts := strings.Split(args[1], ",")
  for _, t := range ts {
    nc.Subscribe(t, func(msg *nats.Msg) {
      log.Printf("Received msg on topic [%s], '%s'\n", t, string(msg.Data))
    })
    log.Printf("Subscribed topic [%s]\n", t)
  }

  runtime.Goexit()
}
