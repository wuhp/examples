package main

import (
    "fmt"
    "log"
    "github.com/shizeeg/gcfg"
)

/*
func ReadFileInto(config interface{}, filename string) error
func ReadInto(config interface{}, reader io.Reader) error
func ReadStringInto(config interface{}, str string) error
*/

func main() {
    cfgStr := `# Comment line 1
; Comment line 2
[section]  # Comment Line 3
multi=value1  # Comment Line 4
multi=value2
x=100`

    cfg := struct {
        Section struct {
            Multi []string
            Var int `gcfg:"x"`
        }
    }{}

    err := gcfg.ReadStringInto(&cfg, cfgStr)
    if err != nil {
        log.Fatalf("Failed to parse gcfg data: %s", err)
    }

    fmt.Println(cfg.Section.Multi)
    fmt.Println(cfg.Section.Var)
}
