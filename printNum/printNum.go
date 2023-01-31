package main

import (
	"log"
	"time"
)

func print(ch chan int) {
	for {
		select {
		case data := <-ch:
			if data > 100 {
				break
			}
			log.Println("recv-data:[%+v]", data)
			ch <- data + 1
		}
	}
}

func main() {
	ch := make(chan int, 1)
	nums := 3
	for i := 0; i < nums; i++ {
		go func() {
			print(ch)
		}()
	}
	ch <- 1

	go hang()
	select {}
}

func hang() {
	time.Sleep(time.Hour)
}
