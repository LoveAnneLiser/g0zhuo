package main

import (
	"fmt"
	"zhuo"
)

func Log(next zhuo.HandlerFunc) zhuo.HandlerFunc {
	return func(ctx *zhuo.Context) {
		fmt.Println("打印请求参数")
		next(ctx)
		fmt.Println("返回执行时间")
	}
}
func main() {

	engine := zhuo.New()
	g := engine.Group("user")
	// 测试前置中间件
	g.Use(func(next zhuo.HandlerFunc) zhuo.HandlerFunc {
		return func(ctx *zhuo.Context) {
			fmt.Println("pre handler")
			next(ctx)
			fmt.Println("post handler")
		}
	})

	g.Get("/hello/get", func(ctx *zhuo.Context) {
		fmt.Println("/hello/get handler")
		fmt.Fprintf(ctx.W, "%s get 欢迎来到卓的Go框架", "zhuo.com")
	}, Log)
	g.Post("/hello/get", func(ctx *zhuo.Context) {
		fmt.Fprintf(ctx.W, "%s post hello", "zhuo.com")
	})

	g.Get("/hello/*/get", func(ctx *zhuo.Context) {
		fmt.Fprintf(ctx.W, "%s post /hello/*/get", "zhuo.com")
	})

	g.Get("/get/:id", func(ctx *zhuo.Context) {
		fmt.Fprintf(ctx.W, "%s get user info path variable", "zhuo.com")
	})
	order := engine.Group("order")
	order.Get("/get/goods", func(ctx *zhuo.Context) {
		fmt.Fprintf(ctx.W, "%s 查询订单", "zhuo.com")
	})
	engine.Run()
}
