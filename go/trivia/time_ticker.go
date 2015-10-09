package main

import (
    "log"
    "time"
    "runtime"
)

func timer1() {
    timer1 := time.NewTicker(1 * time.Second)
    for {
        select {
        case <- timer1.C:
            log.Println("timer 1")
        }
    }
}

func timer2() {
    timer2 := time.NewTicker(2 * time.Second)
    for {
        select {
        case <- timer2.C:
            log.Println("timer 2")
        }
    }
}

func main() {
    go timer1()
    go timer2()

    runtime.Goexit()
}
