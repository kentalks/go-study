package utils

import (
	"fmt"
	"time"
)

type Worker struct {
	Id int
}

func (w *Worker) Process(c chan int) {
	for {
		select {
		case data := <-c:
			fmt.Printf("worker %d got %d\n", w.Id, data)
		case <-time.After(time.Millisecond * 10):
			fmt.Println("Break time")
			time.Sleep(time.Second * 2)
		}
	}
}
