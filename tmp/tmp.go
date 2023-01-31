package main

func main() {
	arrs := []int{1, 2, 0, 3, 0, 4, 0}
	moveNumsZero(arrs)
}

func moveNumsZero(arrs []int) {
	length := len(arrs)
	if length < 1 {
		return
	}
	slow, fast := 0, 0
	for fast < length && slow < length {
		for slow < length && arrs[slow] != 0 {
			slow++
		} //slow 找到第一个0
		fast = slow
		for fast < length && arrs[fast] == 0 {
			fast++
		}
		arrs[slow], arrs[fast] = arrs[fast], arrs[slow]
	}
	return
}
