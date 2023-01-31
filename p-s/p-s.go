package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	consumerNum, productNum := 10, 100
	ch := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(consumerNum + productNum)
	for i := 0; i < consumerNum; i++ {
		go func(idx int) {
			defer wg.Done()
			consume(idx, ch)
		}(i)
	}

	log.Println("start consume! \n start product data!")
	for i := 0; i < productNum; i++ {
		go func(idx int) {
			defer wg.Done()
			product(idx, ch, 1000)
		}(i)
	}
	time.Sleep(time.Second)
	log.Printf("all-done")

	time.Sleep(time.Hour)
}

func product(idx int, ch chan int, limit int) {
	for i := 0; i < limit; i++ {
		ch <- i
		log.Printf("product-idx:[ %d] ,send-data:[%+v]", idx, i)
	}
	log.Printf("product-idx:[ %d] ,product-done", idx)
}

func consume(idx int, ch chan int) {
	for {
		select {
		case data := <-ch:
			log.Printf("consumer-idx:[ %d] ,recv-data:[%+v]", idx, data)
		}
	}
}


