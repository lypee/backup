package main

import (
	"log"
	"sync"
	"time"
)

type PrintNumsEntity struct {
	chs        []chan int
	chanIdxMap map[int]ChanEntity
	recvChan   chan int
}

type ChanEntity struct {
	Index    int
	PrintNum int
	Chan     chan int
}

func Init(cap int) *PrintNumsEntity {
	this := new(PrintNumsEntity)
	this.recvChan = make(chan int, 1)
	this.chs = make([]chan int, 0, cap)
	this.chanIdxMap = make(map[int]ChanEntity, 0)

	for i := 0; i < cap; i++ {
		ch := make(chan int, 1)
		this.chs = append(this.chs, ch)
		this.chanIdxMap[i] = ChanEntity{
			Index:    i,
			PrintNum: i,
			Chan:     ch,
		}
	}

	return this
}

func (this *PrintNumsEntity) PrintNum() {

}

func (this *PrintNumsEntity) CloseChanList() {
	defer func() {
		recover()
	}()
	for _, ch := range this.chs {
		close(ch)
	}
	return
}

const (
	NUMSLIMIT     = 100
	GOROUTINENUMS = 2
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	wg := &sync.WaitGroup{}
	wg.Add(GOROUTINENUMS)

	go hang()

	go func() {
		defer wg.Done()
		print1(ch1, ch2, 1)
	}()

	go func() {
		defer wg.Done()
		print2(ch1, ch2, 2)
	}()
	wg.Wait()

	log.Println("done")
	select {}
}

func print1(ch1, ch2 chan int, index int) {
	for {
		if data, ok := <-ch1; ok {
			if data > NUMSLIMIT {
				return
			}
			log.Printf("index:[%d] recv:[%d]", index, data)
			ch2 <- data + 1
		}
	}
}

func print2(ch1, ch2 chan int, index int) {
	for {
		if data, ok := <-ch2; ok {
			if data > NUMSLIMIT {
				return
			}
			log.Printf("index:[%d] recv:[%d]", index, data)
			ch1 <- data + 1
		}
	}
}

func hang() {
	time.Sleep(time.Duration(10) * time.Second)
}
