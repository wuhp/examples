package main

import (
    "bytes"
    "fmt"
    "text/template"
)

type X struct {
    A *Y
}

type Y struct {
    B int
}

func (y *Y) Fun(a int) int {
    return y.B + a
}

func main() {
    yml := `template is {{.A.Fun 2}}`
    x := new(X)
    x.A = new(Y)
    x.A.B = 1

    var b bytes.Buffer
    t, _ := template.New("").Parse(yml)
    t.Execute(&b, x)
    fmt.Println(b.String())
}
