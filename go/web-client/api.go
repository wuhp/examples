package main

import (
  "net/http"
  "bytes"
  "fmt"
  "encoding/json"
  "io/ioutil"
)

const (
  host = "localhost:8080"
)

///////////////////////////////////////////////////////////////////////////////

type Product struct {
  Id int              `json:"id"`
  Name string         `json:"name"`
  Description string  `json:"description"`
  CreateTs string     `json:"create_ts"`
  LastUpdateTs string `json:"last_udpate_ts"`
}

func ListProduct() []*Product {
  client := &http.Client{}

  req, _:= http.NewRequest(
    "GET",
    fmt.Sprintf("http://%s/api/product", host),
    nil,
  )

  res, _ := client.Do(req)
  defer res.Body.Close()

  var products []*Product
  body, _ := ioutil.ReadAll(res.Body)
  json.Unmarshal(body, &products)

  return products
}

func CreateProduct(name, desc string) int {
  type ainput struct {
    Name string `json:"name"`
    Description string `json:"description"`
  }

  in := new(ainput)
  in.Name = name
  in.Description = desc

  client := &http.Client{}

  inbody, _ := json.Marshal(in)
  req, _:= http.NewRequest(
    "POST",
    fmt.Sprintf("http://%s/api/product", host),
    bytes.NewBuffer(inbody),
  )

  req.Header.Set("Content-Type", "productlication/json")
  res, _ := client.Do(req)
  defer res.Body.Close()

  var m map[string]int
  outbody, _ := ioutil.ReadAll(res.Body)
  json.Unmarshal(outbody, &m)

  return m["id"]
}

func GetProduct(id int) *Product {
  client := &http.Client{}

  req, _:= http.NewRequest(
    "GET",
    fmt.Sprintf("http://%s/api/product/%d", host, id),
    nil,
  )

  res, _ := client.Do(req)
  defer res.Body.Close()

  product := new(Product)
  body, _ := ioutil.ReadAll(res.Body)
  json.Unmarshal(body, product)

  return product
}

func DeleteProduct(id int) {
  client := &http.Client{}

  req, _:= http.NewRequest(
    "DELETE",
    fmt.Sprintf("http://%s/api/product/%d", host, id),
    nil,
  )

  res, _ := client.Do(req)
  defer res.Body.Close()
}
