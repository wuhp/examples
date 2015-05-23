package main

import (
  "io"
  "fmt"
  "os"
  "bufio"
)

func main() {
  rw, err := os.Open("/tmp/abcdxx")
  if err != nil {
    panic(err)
  }
  defer rw.Close()

  rb := bufio.NewReader(rw)
  for {
    line, _, err := rb.ReadLine()
    if err == io.EOF {
      break
    }
    //do something
    fmt.Println(string(line))
  }
}
