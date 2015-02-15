package main

import "os/exec"
import "log"
import "strings"

func main() {
  id, _ := exec.Command("uuidgen").Output()
  log.Printf("==%s==", strings.TrimSuffix(string(id), "\n"))
}
