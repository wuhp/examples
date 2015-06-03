package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "crypto/tls"
)

func main() {
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        },
    }

    req, _ := http.NewRequest(
        "GET",
        "https://api.bitbucket.org/1.0/repositories/daocloud-hwu/argus/src/7cc09113c17f/Dockerfile",
        nil,
    )

    res, err := client.Do(req)
    if err != nil {
        fmt.Println("Fatal error, client do")
        return
    }

    fmt.Println(res.StatusCode)

    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("Fatal error, read body")
    }

    fmt.Println(string(body))
}
