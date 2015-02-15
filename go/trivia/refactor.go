package main

import "fmt"
import "reflect"

type Person struct {
  Name string `tagn`
  Age int `taga`
}

func show(i interface{}) {
  switch i.(type) {
  case *Person:
    t := reflect.TypeOf(i)
    v := reflect.ValueOf(i)
    tag := t.Elem().Field(0).Tag
    fmt.Printf("%v\n", tag)
    name := v.Elem().Field(0).String()
    fmt.Printf("%v\n", name)
  }
}

func main() {
  p := new(Person)
  p.Name = "haipeng"
  p.Age = 30
  show(p)
}
