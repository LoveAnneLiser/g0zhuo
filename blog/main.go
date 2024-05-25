package main

import (
	"fmt"
	"net/http"
	"zhuo"
)

type User struct {
	Name string
}

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
	g.Get("/html", func(ctx *zhuo.Context) {
		ctx.HTML(http.StatusOK, "<h1>小卓学Go</h1>")
	})
	g.Get("/htmlTemplate", func(ctx *zhuo.Context) {
		user := &User{"zhuo"}
		err := ctx.HTMLTemplate("login.html", user, "tpl/login.html", "tpl/header.html")
		if err != nil {
			fmt.Println("err:", err)
		}
	})

	g.Get("/htmlTemplateGlob", func(ctx *zhuo.Context) {
		user := &User{"zhuo"}
		err := ctx.HTMLTemplateGlob("login.html", user, "tpl/*.html")
		if err != nil {
			fmt.Println("err:", err)
		}
	})
	order := engine.Group("order")
	order.Get("/get/goods", func(ctx *zhuo.Context) {
		fmt.Fprintf(ctx.W, "%s 查询订单", "zhuo.com")
	})
	engine.Run()
}
