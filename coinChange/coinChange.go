package main

import (
	"log"
	"sort"
)

func main() {
	coins := []int{186,419,83,408}
	amount := 6249
	log.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	count := 0
	if amount == 0 {
		return count
	}
	sort.Ints(coins)
	for i := len(coins) - 1; i >= 0; {
		if amount == 0 && i < 0 {
			break
		}
		if amount < coins[i] {
			i--
			continue
		}
		amount -= coins[i]
		count++
	}
	if count == 0 || amount > 0{
		return -1
	}
	return count
}
