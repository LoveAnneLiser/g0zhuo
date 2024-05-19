package main

import (
	"fmt"
	"zhuo"
)

func main() {
	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "%s 欢迎来到卓的Go框架", "zhuo.com")
	//})
	//err := http.ListenAndServe(":8111", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	engine := zhuo.New()
	g := engine.Group("user")
	g.Get("/hello", func(ctx *zhuo.Context) {
		fmt.Fprintf(ctx.W, "%s get 欢迎来到卓的Go框架", "zhuo.com")
	})
	g.Post("/hello", func(ctx *zhuo.Context) {
		fmt.Fprintf(ctx.W, "%s post hello", "zhuo.com")
	})
	g.Any("/any", func(ctx *zhuo.Context) {
		fmt.Fprintf(ctx.W, "%s info1", "zhuo.com")
	})
	order := engine.Group("order")
	order.Any("/get", func(ctx *zhuo.Context) {
		fmt.Fprintf(ctx.W, "%s 查询订单", "zhuo.com")
	})
	engine.Run()
}
