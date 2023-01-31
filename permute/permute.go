package main

import "log"

var (
	arrs []int
	pd   []int
	ans  int
)

func main() {
	per()
	log.Println(ans)
}
func per() {
	arrs = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	pd = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	dfs(0)
}

func dfs(n int) {
	if n == 5 {
		ans++
	}
	for idx := 1; idx <= 9; idx++ {
		if pd[idx] == 0 {
			pd[idx] = 1
			arrs[n] = idx
			dfs(n + 1)
			pd[idx] = 0
			arrs[n] = 0
		}
	}
}
