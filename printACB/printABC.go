package main

import (
	"log"
	"time"
)

var (
	nums  = 10
	times = 1
)

func main() {
	chA := make(chan struct{}, 1)
	chB := make(chan struct{}, 1)
	chC := make(chan struct{}, 1)
	go printA(chA, chB)
	go printB(chB, chC)
	go printC(chC, chA)
	chA <- struct{}{}
	time.Sleep(time.Second * 5)
}

func printA(chA, chB chan struct{}) {

	for {
		select {
		case <-chA:
			if times > 30 {
				return
			}
			log.Printf("第 %d 次打印 A", times)
			times++
			chB <- struct{}{}
		}
	}
}

func printB(chB, chC chan struct{}) {

	for {
		select {
		case <-chB:
			if times > 30 {
				return
			}
			log.Printf("第 %d 次打印 B", times)
			times++
			chC <- struct{}{}
		}
	}
}

func printC(chC, chA chan struct{}) {
	for {
		select {
		case <-chC:
			if times > 30 {
				return
			}
			log.Printf("第 %d 次打印 C", times)
			times++
			chA <- struct{}{}
		}
	}
}
