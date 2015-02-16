package main

import (
  "fmt"
  "flag"
  "strconv"
)

func usage() {
  fmt.Println("Usage:")
  fmt.Println("./web-client list_product")
  fmt.Println("./web-client create_product <name> <description>")
  fmt.Println("./web-client show_product <id>")
  fmt.Println("./web-client delete_product <id>")
  fmt.Println()
}

func printProduct(product *Product) {
  fmt.Printf("id: %d\n", product.Id)
  fmt.Printf("name: %s\n", product.Name)
  fmt.Printf("description: %s\n", product.Description)
  fmt.Printf("create_ts: %s\n", product.CreateTs)
  fmt.Printf("last_update_ts: %s\n", product.LastUpdateTs)
}

func listProduct() {
  products := ListProduct()
  for i, product := range(products) {
    fmt.Printf("[ Product %d ]\n", i)
    printProduct(product)
    fmt.Printf("\n")
  }
}

func createProduct(name, desc string) {
  id := CreateProduct(name, desc) 
  fmt.Printf("id: %d\n", id)
}

func showProduct(id int) {
  product := GetProduct(id)
  printProduct(product)
  fmt.Printf("\n")
}

func deleteProduct(id int) {
  DeleteProduct(id)
}

///////////////////////////////////////////////////////////////////////////////

func main() {
  flag.Parse()

  switch flag.Arg(0) {
  case "list_product":
    listProduct()
  case "create_product":
    createProduct(flag.Arg(1), flag.Arg(2))
  case "show_product":
    id, _ := strconv.Atoi(flag.Arg(1))
    showProduct(id)
  case "delete_product":
    id, _ := strconv.Atoi(flag.Arg(1))
    deleteProduct(id)
  default:
    usage()
  }
}
