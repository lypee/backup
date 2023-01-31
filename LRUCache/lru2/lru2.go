package main

func main() {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	return
}


type LRUCache struct {
	Data map[int]int
	Cap  int
	List []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Data: make(map[int]int, 0),
		Cap:  capacity,
		List: make([]int, 0),
	}
}

func (this *LRUCache) Put(key int, value int) {
	oldValue := this.Get(key)
	if oldValue != -1 {
		this.Data[key] = value
		return
	} else { // 原值不存在
		// 比较cap
		if len(this.List) >= this.Cap {
			// 淘汰策略
			this.removeLast()
		}
		//新增策略
		this.Data[key] = value
		this.moveToFirst(key)
		return
	}
	return
}

func (this *LRUCache) Get(key int) int {
	if data, exist := this.Data[key]; exist {
		this.moveToFirst(key)
		return data
	}
	return -1
}

func (this *LRUCache) moveToFirst(key int) {
	newList := make([]int, 0, len(this.List))

	for index := range this.List {
		if this.List[index] == key {
			newList = append(newList, this.List[index])
			break
		}
	}
	if len(newList) == 0 {
		newList = append(newList, key)
	}
	for idx := range this.List {
		if this.List[idx] == key {
			continue
		}
		newList = append(newList, this.List[idx])
	}
	this.List = newList
	return
}

// removeLast 淘汰最久未使用
func (this *LRUCache) removeLast() int {
	oldKey := this.List[len(this.List)-1]
	delete(this.Data, oldKey)
	this.List = this.List[:len(this.List)-1]
	return oldKey
}
