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
	engine.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s 欢迎来到卓的Go框架", "zhuo.com")
	})
	engine.Run()
}
