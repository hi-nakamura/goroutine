package main

import (
    "./count"
)

func main() {
    done := make(chan bool)

    go count.CountDown("name_1", 10, done)
    go count.CountDown("name_2", 10, done)

    <-done
    <-done
}