package zhuo

import (
	"fmt"
	"log"
	"net/http"
)

const ANY = "ANY"

type HandleFunc func(ctx *Context)

// 分组路由
type routerGroup struct {
	name             string
	handleFuncMap    map[string]map[string]HandleFunc
	handlerMethodMap map[string][]string
}

//func (r *routerGroup) Add(name string, handleFunc HandleFunc) {
//	r.handleFuncMap[name] = handleFunc
//}

// 提取操作 减少冗余
func (r *routerGroup) handle(name string, method string, handleFunc HandleFunc) {
	_, ok := r.handleFuncMap[name]
	if !ok {
		r.handleFuncMap[name] = make(map[string]HandleFunc)
	}

	_, ok = r.handleFuncMap[name][method]
	if ok { // 重复定义了handleFunc
		panic(method + ": 同一个路由，不能重复")
	}
	r.handleFuncMap[name][method] = handleFunc
	r.handlerMethodMap[method] = append(r.handlerMethodMap[method], name)
}

func (r *routerGroup) Any(name string, handleFunc HandleFunc) {
	r.handle(name, ANY, handleFunc)
}

func (r *routerGroup) Get(name string, handleFunc HandleFunc) {
	r.handle(name, http.MethodGet, handleFunc)
}

func (r *routerGroup) Post(name string, handleFunc HandleFunc) {
	r.handle(name, http.MethodPost, handleFunc)
}

// user get-> handle
// goods
// order
type router struct {
	routerGroup []*routerGroup // 分组
}

// Group 实现分组
func (r *router) Group(name string) *routerGroup {
	routerGroup := &routerGroup{
		name:             name,
		handleFuncMap:    make(map[string]map[string]HandleFunc),
		handlerMethodMap: make(map[string][]string),
	}
	r.routerGroup = append(r.routerGroup, routerGroup)
	return routerGroup
}

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router: router{},
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	// fmt.Println(method) 得到方法
	for _, group := range e.routerGroup {
		for name, methodHandle := range group.handleFuncMap {
			url := "/" + group.name + name
			log.Println("url: ", url)
			if r.RequestURI == url { // 路由匹配
				ctx := &Context{
					W: w,
					R: r,
				}
				handle, ok := methodHandle[ANY]
				if ok {
					handle(ctx)
					return
				}
				handle, ok = methodHandle[method]
				if ok {
					handle(ctx)
					return
				}
				w.WriteHeader(http.StatusMethodNotAllowed) // 405
				fmt.Fprintln(w, r.RequestURI+" "+method+" 这个请求方式不允许")
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound) // 404
	fmt.Fprintf(w, "%s not found \n", r.RequestURI)
}

func (e *Engine) Run() {
	// user key :get value: func
	//for _, group := range e.routerGroup {
	//	for key, value := range group.handleFuncMap {
	//		http.HandleFunc("/"+group.name+key, value)
	//		log.Printf("/" + group.name + key)
	//	}
	//}
	http.Handle("/", e)
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
