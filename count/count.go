package count

import (
    "fmt"
    "time"
)

type Count struct {
    name   string
    count  int
}

func (this *Count) CountDown() {
    for i := this.count; i >= 0; i-- {
        fmt.Printf("name : %s count : %d\n", this.name, i)
        time.Sleep(1 * time.Second) // 1秒Sleep
    }
}

func CountDown(name string, start_count int, done chan bool) {
    count := &Count{name, start_count}
    count.CountDown()
    done <- true
}
