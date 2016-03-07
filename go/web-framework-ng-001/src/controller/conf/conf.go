package conf

import (
    "os"
    "errors"
)

var DbConnectionUrl string

func Initialize() error {
    DbConnectionUrl = os.Getenv("DB_CONNECTION_URL")
    if len(DbConnectionUrl) == 0 {
        return errors.New("ERROR: empty env var DB_CONNECTION_URL")
    }

    return nil
}
