package main

import (
    "os"
    "flag"
    "log"
    "time"

    "github.com/VividCortex/godaemon"
)

func main() {
    path := flag.String("o", "/tmp/daemon.log", "log file")
    flag.Parse()

    f, e := os.OpenFile(*path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if e != nil {
        log.Fatalf("ERROR: Failed to open log file `%s`, with err `%s`\n", *path, e.Error())
    }

    godaemon.MakeDaemon(&godaemon.DaemonAttr{})

    log.SetOutput(f)

    for {
        log.Println("abc")
        time.Sleep(1 * time.Second)
    }
}

