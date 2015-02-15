package main

import (
  "time"
)

var c chan int

func ready() {
  <-c
}

func main() {
  c = make(chan int)
  go ready()
  go ready()
  go ready()
  go ready()
  go ready()
  time.Sleep(50 * time.Second)
  c <- 1
}
