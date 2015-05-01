package main

import (
    "fmt"
    "log"
    "gopkg.in/yaml.v2"
)

var data = `
build_image: ubuntu:12.04!

services: [mysql, redis]

env:
  - STR1 = "11"
  - STR2 = "22"

script:
- echo "this is script segment"
- echo "this is script segment"

mapping:
  key1: value1
  key2: [value21, value22]
  key3:
    - value31
    - value32
`

type Part struct {
    Keya string   `yaml:"key1"`
    Keyb []string `yaml:"key2"`
    Keyc []string `yaml:"key3"`
}

type Yml struct {
    BuildImage   string   `yaml:"build_image,omitempty"`
    Services     []string `yaml:"services,omitempty"`
    Env          []string `yaml:"env,omitempty"`
    Script       []string `yaml:"script,omitempty"`
    Mapping      Part     `yaml:"mapping,omitempty"`
}

func main() {
    t := Yml{}

    err := yaml.Unmarshal([]byte(data), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    fmt.Printf("--- t:\n%v\n\n", t)

    d, err := yaml.Marshal(&t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    fmt.Printf("--- t dump:\n%s\n\n", string(d))

    m := make(map[interface{}]interface{})

    err = yaml.Unmarshal([]byte(data), &m)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    fmt.Printf("--- m:\n%v\n\n", m)

    d, err = yaml.Marshal(&m)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
