package main

import (
  "fmt"
  "time"
  "os"
  "os/signal"
)

func main() {
  go signalListen()
  time.Sleep(time.Hour)
}

func signalListen() {
  c := make(chan os.Signal)
  signal.Notify(c)
  for {
    s := <-c
    fmt.Println("Get signal:", s)
    os.Exit(1)
  }
}
