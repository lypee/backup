package interview

import "sync"

type Node struct {
	Key string
	Val int
}

type LRUCache struct {
	DataMap map[string]Node
	List    []Node
	Cap     int
	mu      sync.Mutex
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		DataMap: make(map[string]Node, 0),
		List:    make([]Node, 0, cap),
		Cap:     cap,
	}
}

func (srv *LRUCache) Get(key string) (Node, int, bool) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	if len(srv.List) < 1 {
		// error
		return Node{}, 0, false
	}
	node, exist := srv.DataMap[key]
	if !exist {
		// return error
		return Node{}, 0, false
	}
	idx := 0
	for idx = 0; idx < len(srv.List); idx++ {
		if srv.List[idx] == node {
			break
		}
	}
	srv.List = append(srv.List[:idx], srv.List[idx:len(srv.List)]...)
	srv.List = append(srv.List, node)
	return node, idx, true
}

func (srv *LRUCache) Set(key string, val int) {
	srv.mu.Lock()
	defer srv.mu.Unlock()

	newNode := Node{Key: key, Val: val}
	oldNode, idx, exist := srv.Get(key)
	// 是否已满？
	if len(srv.List) < srv.Cap {
		// 未满
		// 存在
		if exist {
			srv.Update(newNode, oldNode, idx)
			return
		} else {
			srv.Add(newNode)
		}
	} else {
		// 已满
		// 存在
		if exist {
			srv.Update(newNode, oldNode, idx)
			return
		}
		// 不存在
		disuseNode := srv.List[len(srv.List)-1]
		delete(srv.DataMap, disuseNode.Key)
		srv.Add(newNode)
		srv.List = srv.List[1:len(srv.List)]
	}
}

// Add -默认不存在
func (srv *LRUCache) Add(node Node) {
	srv.List = append([]Node{node}, srv.List...)
	srv.DataMap[node.Key] = node
	return
}

// Update 默认存在
func (srv *LRUCache) Update(node, oldNode Node, idx int) {
	oldNode.Val = node.Val
	srv.DataMap[node.Key] = oldNode
	list := append(srv.List[:idx], srv.List[idx:len(srv.List)]...)
	list = append([]Node{node}, list...)
	srv.List = list
	return
}
