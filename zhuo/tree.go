package zhuo

import "strings"

// 前缀树
type treeNode struct {
	name       string
	children   []*treeNode
	routerName string
}

// Put put path: /user/get/:id
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
			node := &treeNode{
				name:     name,
				children: make([]*treeNode, 0),
			}
			children = append(children, node)
			t.children = children
			t = node
		}
	}
	t = root
}

// Get get path: /user/get/1
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
					return node
				}
			}
		}
	}
	return nil
}
