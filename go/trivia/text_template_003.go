package main

import (
    "bytes"
    "fmt"
    "text/template"
)

type X struct {
    A []*Y
}

func (x *X) Funx(a int) []*Y {
    x.A = make([]*Y, a)
    for i := 0; i < a; i++ {
        x.A[i] = new(Y)
        x.A[i].B = i
    }

    return x.A
}

type Y struct {
    B int
}

func (y *Y) Funy(a int) int {
    return y.B + a
}

func main() {
    yml := `a{{range $x := .Funx 3}}x{{$x.Funy 10}}y{{end}}b`
    x := new(X)

    var b bytes.Buffer
    t, _ := template.New("").Parse(yml)
    t.Execute(&b, x)
    fmt.Println(b.String())
}
