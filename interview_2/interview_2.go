package main

import (
	"fmt"
)

/*
题目描述：扁平数组转Tree

给定一个扁平数组，数组内每个对象的id属性是唯一的。每个对象具有pid属性，pid属性为0表示为根节点（根节点只有一个），其它表示自己的父节点id。
编写一段程序，输入为给定的扁平数组，输出要求为一个树结构，为其中每个对象增加children数组属性（里面存放child对象）。

解法有很多种，性能最优的方案最佳。可以不用处理输入输出，专注实现核心逻辑即可

输入：
[
  {"id": 1, "name": "部门1", "pid": 0},
  {"id": 2, "name": "部门2", "pid": 1},
  {"id": 3, "name": "部门3", "pid": 1},
  {"id": 4, "name": "部门4", "pid": 3},
  {"id": 5, "name": "部门5", "pid": 4},
  {"id": 6, "name": "部门6", "pid": 3}
]
type DepNode struct{
	Id int
	Name string
	PId int
	Children []DepNode
}
输出：
{
  "id": 1,
  "name": "部门1",
  "pid": 0,
  "children": [
    {
      "id": 2,
      "name": "部门2",
      "pid": 1,
      "children": []
    },
    {
      "id": 3,
      "name": "部门3",
      "pid": 1,
      "children": [
        {
          "id": 4,
          "name": "部门4",
          "pid": 3,
          "children": [
            {
              "id": 5,
              "name": "部门5",
              "pid": 4,
              "children": []
            }
          ]
        },
        {
          "id": 6,
          "name": "部门6",
          "pid": 3,
          "children": []
        }
      ]
    }
  ]
}
*/

func main() {
	str := `[
  {"id": 1, "name": "部门1", "pid": 0},
  {"id": 2, "name": "部门2", "pid": 1},
  {"id": 3, "name": "部门3", "pid": 1},
  {"id": 4, "name": "部门4", "pid": 3},
  {"id": 5, "name": "部门5", "pid": 4},
  {"id": 6, "name": "部门6", "pid": 3}
]`
	fmt.Println(str)
}

type DepartmentSrv struct {
	DepInfo *DepNode // 指向 根部门节点
}

type DepNode struct {
	Id       int
	PId      int
	Name     string
	Children []*DepNode
}

type TreeNode struct {
	Id   int
	Name string
	PId  int
}

func (srv *DepartmentSrv) Add(node *DepNode) (bool, error) {
	fatherNode := srv.Search(srv.DepInfo, node.PId)
	if fatherNode == nil {
		depNode := &DepNode{
			Id:       node.Id,
			PId:      0,
			Name:     node.Name,
			Children: make([]*DepNode, 0),
		}
		srv.DepInfo = depNode
		return true, nil
	}
	fatherNode.Children = append(fatherNode.Children, node)
	return true, nil
}

func (srv *DepartmentSrv) Search(node *DepNode, id int) *DepNode {
	if node == nil {
		return nil
	}
	if node.Id == id {
		return node
	}
	stack := []*DepNode{}
	stack = append(stack, node)
	for len(stack) != 0 {
		n := len(stack)
		for i := 0; i < n; i++ {
			subNode := stack[0]
			if subNode.Id == id {
				return subNode
			}
			stack = stack[1:len(stack)]
			for j := 0; j < len(subNode.Children); j++ {
				stack = append(stack, subNode.Children[j])
			}
		}
	}
	return nil
}
