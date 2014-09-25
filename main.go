package main

func main() {
    done := make(chan bool)

    go func() {
        CountDown("name_1", 10)
        done <- true
    }()

    go func() {
        CountDown("name_2", 10)
        done <- true
    }()

    <-done
    <-done
}
