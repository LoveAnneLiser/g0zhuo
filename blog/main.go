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
	g.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s 欢迎来到卓的Go框架", "zhuo.com")
	})
	g.Post("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s hello", "zhuo.com")
	})
	g.Any("/any", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s info1", "zhuo.com")
	})
	order := engine.Group("order")
	order.Add("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s 查询订单", "zhuo.com")
	})
	engine.Run()
}
