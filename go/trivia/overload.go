package main


func f() {
}

func f(i int) {
  i += 1
}

func main() {
  f()
  f(1)
}
