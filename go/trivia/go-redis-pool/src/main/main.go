package main

import (
  "fmt"
  "time"
  "github.com/garyburd/redigo/redis"
)

// Single redis connection can not be paralleled, so use
// connect pool in paralleled environment

func main() {

  pool := &redis.Pool{
    MaxIdle: 3,
    IdleTimeout: 60 * time.Second,
    Dial: func () (redis.Conn, error) {
      c, err := redis.Dial("tcp", "localhost:6379")
      if err != nil {
        return nil, err
      }
/*
      if _, err := c.Do("AUTH", "password"); err != nil {
         c.Close()
         return nil, err
      }
*/
      return c, err
    },
  }

  // ping
  conn := pool.Get()
  _, err := conn.Do("PING")
  if err != nil {
    panic("Fail to connect to redis server")
  }
  conn.Close()

fmt.Println("--------")
fmt.Printf("%+v\n", pool)
fmt.Println("--------")

  conn = pool.Get()
  conn.Send("MULTI")
  conn.Send("RPUSH", "array", 1, 2, "a")
  conn.Send("SET", "number", 3)
  _, err = conn.Do("EXEC")
  if err != nil {
    fmt.Println("Failed")
  }
  conn.Close()

  conn = pool.Get()
  defer conn.Close()

  conn.Send("MULTI")
  conn.Send("EXISTS", "number")
  conn.Send("LRANGE", "array", 0, -1)
  values, err := redis.Values(conn.Do("EXEC"))
  if err != nil {
    fmt.Println("Failed")
  }

fmt.Println("--------")
fmt.Printf("%+v\n", values)
fmt.Println("--------")

  var exist int
  var elements []string

  _, err = redis.Scan(values, &exist, &elements)
  if err != nil {
    fmt.Println("Failed")
  }

  fmt.Println(exist)
  fmt.Println(elements)
}
