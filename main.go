package main

import (
    "./count"
)

func main() {
    done := make(chan bool)

    go func() {
        count.CountDown("name_1", 10)
        done <- true
    }()

    go func() {
        count.CountDown("name_2", 10)
        done <- true
    }()

    <-done
    <-done
}