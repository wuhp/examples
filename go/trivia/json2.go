package main

import "encoding/json"
import "fmt"

func main() {
    x := make([]int, 0)
    y, _ := json.Marshal(x)
    fmt.Println(string(y))

    x = append(x, 1)
    z, _ := json.Marshal(x)
    fmt.Println(string(z))

    m, _ := json.Marshal(map[string]int64{"id":12})
    fmt.Println(string(m))
}
