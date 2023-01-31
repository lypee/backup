package main

import (
	"fmt"
	"sort"
)

var list []int

func main() {
	list = make([]int, 0)
	getNum()
	sort.Ints(list)
	var idx int
	for {
		_, err := fmt.Scanf("%d", &idx)
		if err != nil {
			fmt.Println("err: ", err)
			break
		}
		if idx < 1 || idx > len(list) {
			continue
		}
		fmt.Println(list[idx-1])
	}
}

func getNum() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			for k := 0; k < 20; k++ {
				cur := flow(2, i) * flow(3, j) * flow(7, k)
				if cur < 0 {
					continue
				}
				list = append(list, cur)
			}
		}
	}
}

func flow(a, b int) int {
	cnt := 1
	for b != 0 {
		if b&1 != 0 {
			cnt *= a
		}
		b >>= 1
		a *= a
	}
	return cnt
}
