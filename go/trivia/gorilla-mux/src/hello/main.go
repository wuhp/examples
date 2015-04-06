package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
}

func Greet(w http.ResponseWriter, r *http.Request) {
  name := mux.Vars(r)["name"]
  w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
}

func ProcessPathVariables(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  name := vars["name"]
  job := vars["job"]
  age := vars["age"]
  w.Write([]byte(fmt.Sprintf("Name is %s ", name)))
  w.Write([]byte(fmt.Sprintf("Job is %s ", job)))
  w.Write([]byte(fmt.Sprintf("Age is %s ", age)))
}

func main() {
  mx := mux.NewRouter()

  mx.HandleFunc("/", SayHelloWorld)
  mx.HandleFunc("/{name}", Greet)
  mx.HandleFunc("/person/{name}/{job}/{age:[0-9]+}", ProcessPathVariables)

  http.ListenAndServe(":8080", mx)
}
