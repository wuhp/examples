package main

import "runtime"

func main() {
  go func() {
    panic("xx")
  }()
  runtime.Goexit()
}
