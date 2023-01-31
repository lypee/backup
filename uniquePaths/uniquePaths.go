package main

import "log"



func main() {
	m, n := 3, 2
	log.Println(uniquePaths(m, n))
}

var (
	path [][]int
)

func uniquePaths(m int, n int) int {
	path = make([][]int, m)
	for i := range path {
		path[i] = make([]int, n)
	}
	path[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				path[i][j] = 1
			} else {
				path[i][j] = path[i-1][j] + path[i][j-1]
			}
		}
	}
	return path[m-1][n-1]
}
