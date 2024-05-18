package main

import (
	"fmt"
	"net/http"
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
	g.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s 欢迎来到卓的Go框架", "zhuo.com")
	})
	g.Add("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s info信息", "zhuo.com")
	})
	order := engine.Group("order")
	order.Add("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s 查询订单", "zhuo.com")
	})
	engine.Run()
}
