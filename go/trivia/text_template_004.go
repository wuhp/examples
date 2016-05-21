package main

import (
    "bytes"
    "fmt"
    "text/template"
)

type Deployment struct {
    S []*Service
}

func (d *Deployment) Service(name string) *Service {
    return d.S[0]
}

type Service struct {
    I []*Instance
}

func (s *Service) Instances() []*Instance {
    return s.I
}

type Instance struct {
    N *Node
}

func (i *Instance) Node() *Node {
    return i.N
}

type Node struct {
    N *Network
}

func (n *Node) Network(name string) *Network {
    return n.N
}

type Network struct {
    Ip string
}

func (n *Network) Address() string {
    return n.Ip
}

func (d *Deployment) init(instNum int) {
    d.S = make([]*Service, 1)
    d.S[0] = new(Service)
    d.S[0].I = make([]*Instance, instNum)
    for i := 0; i < instNum; i++ {
        d.S[0].I[i] = new(Instance)
        d.S[0].I[i].N = new(Node)
        d.S[0].I[i].N.N = new(Network)
        d.S[0].I[i].N.N.Ip = "192.168.1."+fmt.Sprintf("%d", i+1)
    }
}

func main() {
    d := new(Deployment)
    d.init(3)
    yml := `{{range $instance := (.Service "web").Instances}}
           server {{($instance.Node.Network "inner").Address}}
           {{end}}`
 
    var b bytes.Buffer
    t, _ := template.New("").Parse(yml)
    t.Execute(&b, d)
    fmt.Println(b.String())
}
