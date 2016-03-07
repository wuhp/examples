package model

import (
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Initialize(uri string) error {
    var err error 
    if db, err = sql.Open("mysql", uri); err != nil {
        return err
    }

    return db.Ping()
}
