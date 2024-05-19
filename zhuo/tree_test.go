package zhuo

import (
	"fmt"
	"testing"
)

func TestTreeNode_Put(t *testing.T) {
	root := &treeNode{name: "/", children: make([]*treeNode, 0)}
	root.Put("/user/get/:id")
	root.Put("/user/create/hello")
	root.Put("/user/create/aaa")
	root.Put("/order/create/hello")

	node := root.Get("/user/get/1")
	fmt.Println(node)

	node = root.Get("/user/create/hello")
	fmt.Println(node)

	node = root.Get("/user/create/aaa")
	fmt.Println(node)

	node = root.Get("/order/create/hello")
	fmt.Println(node)

	node = root.Get("/order/hello")
	fmt.Println(node)
}
