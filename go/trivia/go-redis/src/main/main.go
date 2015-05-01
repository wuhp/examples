package main

import (
  "fmt"
  "github.com/garyburd/redigo/redis"
)

func main() {
  c, err := redis.Dial("tcp", "localhost:6379")
  if err != nil {
    panic(err)
  }

  defer c.Close()

  c.Do("SET", "k1", 1)

  n, _ := redis.Int(c.Do("GET", "k1"))
  fmt.Printf("%#v\n", n)

  n, _ = redis.Int(c.Do("INCR", "k1"))
  fmt.Printf("%#v\n", n)
}
