package main

import (
    "fmt"
    "log"
    "gopkg.in/yaml.v2"
)

type Data struct {
    Content string `yaml:"content"`
}

func printYml(data string) {
    d := Data{}
    err := yaml.Unmarshal([]byte(data), &d)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    fmt.Println("start")
    fmt.Println(d.Content)
    fmt.Println("done")
}

func main() {
    data := `
content: >
  1:11
  222
  333
`

    printYml(data)

    data = `
content: |
  1:11
  222
  333`

    printYml(data)
}
