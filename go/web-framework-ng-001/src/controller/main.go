package main

import (
    "log"
    "net/http"
    "runtime"

    "controller/conf"
    "controller/handler"
    "controller/migrate"
    "controller/model"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    if err := conf.Initialize(); err != nil {
        log.Fatalf("ERROR: init configuration failed, %s\n", err.Error())
    }

    if err := model.Initialize(conf.DbConnectionUrl); err != nil {
        log.Fatalf("ERROR: init db failed, %s\n", err.Error())
    }

    if err := migrate.Run(conf.DbConnectionUrl); err != nil {
        log.Fatalf("ERROR: schema migration failed, %s\n", err.Error())
    }

    log.Printf("Listening on 0.0.0.0:9080 ...\n")
    log.Fatal(http.ListenAndServe(":9080", handler.NewRouter()))
}
