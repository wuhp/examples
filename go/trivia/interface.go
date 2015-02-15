package main

import "fmt"

///////////////////////////////////

type S struct {
  i int
}

func (s *S) Get() int {
  return s.i
}

func (s *S) Put(v int) {
  s.i = v
}

////////////////////////////////////

type R struct {
  i int
}

func (r *R) Get() int {
  return r.i
}

func (r *R) Put(v int) {
  r.i = v
}

////////////////////////////////////

type M struct {
  i int
}

///////////////////////////////////

type I interface {
  Get() int
  Put(int)
}

///////////////////////////////////

func f1(o interface{}) {
  if t, ok := o.(I); ok {
    fmt.Printf("%v\n", t)
    return
  }
  fmt.Printf("Invalid type.\n")
}

func f2(o interface{}) {
  a := o.(I).Get()
  fmt.Printf("%d\n", a)
}

func main() {
  var r R
//  r := new(R)
//  s := new(S)
//  m := new(M)
//  f1(r)
//  f1(s)
//  f1(m)
  f2(r)
}
