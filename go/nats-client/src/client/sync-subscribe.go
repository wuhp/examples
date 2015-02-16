package main

import (
  "log"
  "os"
  "fmt"
  "strings"
  "runtime"

  "github.com/apcera/nats"
)

func main() {
  args := os.Args
  if len(args) != 2 {
    fmt.Println(args[0], "<topic>")
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

  topic := args[1]
  nc.Subscribe(topic, func(msg *nats.Msg) {
    log.Printf("Received msg on topic [%s], '%s'\n", topic, string(msg.Data))
    res := "Sync response to " + topic
    nc.Publish(msg.Reply, []byte(res))
  })
  log.Printf("Subscribed topic [%s]\n", topic)

  runtime.Goexit()
}
