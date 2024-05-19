package zhuo

import "strings"

// 前缀树
type treeNode struct {
	name       string
	children   []*treeNode
	routerName string
	isEnd      bool
}

// Put put path: /user/get/:id
// bug2: 当输入路径/user/hello 但是路径只有user/hello/get时，应该报404错
// 但是没有报错，所以我们检查Put函数，添加这个IsEnd变量
func (t *treeNode) Put(path string) {
	root := t
	strs := strings.Split(path, "/")
	for index, name := range strs {
		if index == 0 {
			continue
		}
		children := t.children
		isMatch := false
		for _, node := range children {
			if node.name == name {
				isMatch = true
				t = node
				break
			}
		}
		if !isMatch { // 没有匹配上
			isEnd := false
			if index == len(strs)-1 {
				isEnd = true
			}
			node := &treeNode{name: name, children: make([]*treeNode, 0), isEnd: isEnd}
			children = append(children, node)
			t.children = children
			t = node
		}
	}
	t = root
}

// Get get path: /user/get/1
// /hello

func (t *treeNode) Get(path string) *treeNode {
	strs := strings.Split(path, "/")
	routerName := ""
	for index, name := range strs {
		if index == 0 {
			continue
		}
		children := t.children
		isMatch := false
		for _, node := range children {
			if node.name == name || node.name == "*" || strings.Contains(node.name, ":") {
				isMatch = true
				routerName += "/" + node.name
				node.routerName = routerName
				t = node
				if index == len(strs)-1 {
					return node
				}
				break
			}
		}
		if !isMatch {
			for _, node := range children {
				// /user/**
				// user/get/userInfo
				// user/aa/bb
				if node.name == "**" {
					routerName += "/" + node.name
					node.routerName = routerName
					return node
				}
			}
		}
	}
	return nil
}
