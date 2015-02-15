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

    var s []string
    s = append(s, "aaa")
    s = append(s, "bbb")
    x := new(Response2)
    x.Page = 1
    x.Fruits = s
    y, _ := json.Marshal(*x)
    fmt.Println(string(y))
}
