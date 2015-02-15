package main
import "encoding/json"
import "fmt"

type Response2 struct {
    Page   int   `json:"page"`
    Fruit string `json:"fruits"`
}

func main() {
    str := `{"page":0, "fruits":"apple", "aaa":"bbb"}`
    var res Response2
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
}
