package main

import (
    "log"
    "gopkg.in/yaml.v2"
)

var data = `
mappings:
- key: key1
  log: /tmp/1
- key: key2
  log: /tmp/2
`

type Part struct {
    Key string `yaml:"key"`
    Log string `yaml:"log"`
}

type Yml struct {
    Mappings     []Part     `yaml:"mappings"`
}

func main() {
    t := Yml{}

    err := yaml.Unmarshal([]byte(data), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    log.Printf("--- t:\n%v\n\n", t)
}
