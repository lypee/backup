package main

import (
	"log"
	"sync"
	"time"
)

const (
	NUMSLIMIT     = 100
	GOROUTINENUMS = 2
)

type PrintNumsEntity struct {
	ch []chan int
}

func main() {
	ch := make(chan int, 1)
	ch <- 1
	wg := &sync.WaitGroup{}
	wg.Add(GOROUTINENUMS)

	go hang()

	for i := 0; i < GOROUTINENUMS; i++ {
		go func() {
			defer wg.Done()
			printNums(ch, NUMSLIMIT)
		}()
	}
	wg.Wait()
	log.Println("done")
	select {}
}

func printNums(ch chan int, limit int) {
	for {
		if data, ok := <-ch; ok {
			if data > limit {
				return
			}
			log.Printf("recv:[%d]", data)
			ch <- data + 1
		}
	}
}

func hang() {
	time.Sleep(time.Duration(10) * time.Second)
}
