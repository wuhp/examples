package main

import "fmt"
import "regexp"

func main() {
  pattern := `/api/app/(\d+)$`
  re := regexp.MustCompile(pattern)
  match, _ := regexp.MatchString(pattern, `/api/app/12`)
  if match {
    fmt.Println(re.FindStringSubmatch(`/api/app/12`))
  }
}
