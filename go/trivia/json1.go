package main
import "encoding/json"
import "fmt"

type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
    str := `[{"page":0, "fruits": ["apple", "peach"]}]`
    var res []Response2
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)

    var x []Response2
    var s []string
    s1 := append(s, "aaa")
    s2 := append(s, "aaa")
    x = append(x, Response2{1, s1})
    x = append(x, Response2{2, s2})
    y, _ := json.Marshal(x)
    fmt.Println(string(y))
}
