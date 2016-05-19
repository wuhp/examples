package main

import (
    "bytes"
    "fmt"
    "text/template"
)

type X struct {
    A *Y
}

func (x *X) Funx(a int) *Y {
    return x.A
}

type Y struct {
    B int
}

func (y *Y) Funy(a int) int {
    return y.B + a
}

func main() {
    yml := `template is {{with .Funx 2}}{{.Funy 2}}{{end}}`
    x := new(X)
    x.A = new(Y)
    x.A.B = 1

    var b bytes.Buffer
    t, _ := template.New("").Parse(yml)
    t.Execute(&b, x)
    fmt.Println(b.String())
}
