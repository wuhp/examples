package main

import (
  "net/http"
  "flag"
  "encoding/json"
  "io/ioutil"
  "log"
  "os"
  "fmt"
)

type Config struct {
  Database Storage `json:"database"`
  Port int         `json:"port"`
  Log string       `json:"log"`
}

type Storage struct {
  Db string     `json:"db"`
  Port int      `json:"port"`
  Host string   `json:"host"`
  User string   `json:"user"`
  Passwd string `json:"passwd"`
}

func parseArgs() *Config {
  c := flag.String("c", "/tmp/controller.json", "config file")
  flag.Parse()

  bs, e := ioutil.ReadFile(*c)
  if e != nil {
    log.Fatalln(e)
  }

  config := new(Config)
  e = json.Unmarshal(bs, config)
  if e != nil {
    log.Fatalln(e)
  }

  return config
}

func initLog(path string) {
  if path == "" {
    return
  }

  f, e := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if e != nil {
    log.Fatalln(e)
  }

  log.SetOutput(f)
}

func main() {
  config := parseArgs()
  initLog(config.Log)

  if err := ConnectDB(
    config.Database.Host,
    config.Database.Port,
    config.Database.User,
    config.Database.Passwd,
    config.Database.Db,
  ); err != nil {
    log.Fatalln(err)
  }

  http.HandleFunc("/api/", Api)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}


