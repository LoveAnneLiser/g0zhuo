package zhuo

import (
	"fmt"
	"log"
	"net/http"
)

const ANY = "ANY"

type HandlerFunc func(ctx *Context)

// 定义中间件
type MiddlewareFunc func(handlerFunc HandlerFunc) HandlerFunc

// 分组路由
type routerGroup struct {
	name             string
	handlerFuncMap   map[string]map[string]HandlerFunc
	handlerMethodMap map[string][]string
	treeNode         *treeNode
	middlewares      []MiddlewareFunc // 中间件
}

func (r *routerGroup) Use(middlewareFunc ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, middlewareFunc...)
}

func (r *routerGroup) methodHandle(h HandlerFunc, ctx *Context) {
	// 中间件
	if r.middlewares != nil {
		for _, middlewareFunc := range r.middlewares {
			h = middlewareFunc(h)
		}
	}
	h(ctx)
}

//func (r *routerGroup) Add(name string, handlerFunc HandlerFunc) {
//	r.handlerFuncMap[name] = handlerFunc
//}

// 提取操作 减少冗余
func (r *routerGroup) handle(name string, method string, handlerFunc HandlerFunc) {
	_, ok := r.handlerFuncMap[name]
	if !ok {
		r.handlerFuncMap[name] = make(map[string]HandlerFunc)
	}

	_, ok = r.handlerFuncMap[name][method]
	if ok { // 重复定义了handlerFunc
		panic(method + ": 同一个路由，不能重复")
	}
	r.handlerFuncMap[name][method] = handlerFunc
	r.treeNode.Put(name)
	//r.handlerMethodMap[method] = append(r.handlerMethodMap[method], name)
}

func (r *routerGroup) Any(name string, handlerFunc HandlerFunc) {
	r.handle(name, ANY, handlerFunc)
}

func (r *routerGroup) Get(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodGet, handlerFunc)
}

func (r *routerGroup) Post(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodPost, handlerFunc)
}

func (r *routerGroup) Delete(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodDelete, handlerFunc)
}
func (r *routerGroup) Put(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodPut, handlerFunc)
}
func (r *routerGroup) Patch(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodPatch, handlerFunc)
}
func (r *routerGroup) Options(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodOptions, handlerFunc)
}
func (r *routerGroup) Head(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodHead, handlerFunc)
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
		handlerFuncMap:   make(map[string]map[string]HandlerFunc),
		handlerMethodMap: make(map[string][]string),
		treeNode:         &treeNode{name: "/", children: make([]*treeNode, 0)},
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
	e.httpRequestHandle(w, r)
}

func (e *Engine) httpRequestHandle(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	// fmt.Println(method) 得到方法
	for _, group := range e.routerGroup {
		routerName := SubStringLast(r.RequestURI, "/"+group.name)
		//fmt.Printf("截掉后的路由名称：%s\n", routerName)
		// /get/1
		node := group.treeNode.Get(routerName)
		if node == nil || !node.isEnd {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, r.RequestURI+" not found")
			return
		}
		//fmt.Printf("node名称：%s\n", node.name) // bug1: 要先判断node是否为空，如果直接打印为空的值 会直接报错。（空指针）
		if node != nil {
			// 路由匹配
			ctx := &Context{
				W: w,
				R: r,
			}
			handle, ok := group.handlerFuncMap[node.routerName][ANY]
			//			log.Println("node.routerName名称：", node.routerName)
			if ok {
				group.methodHandle(handle, ctx)
				return
			}
			handle, ok = group.handlerFuncMap[node.routerName][method]
			if ok {
				group.methodHandle(handle, ctx)
				return
			}
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			fmt.Fprintln(w, r.RequestURI+" "+method+" 这个请求方式不允许")
			return

		}
	}
	w.WriteHeader(http.StatusNotFound) // 404
	fmt.Fprintf(w, "%s not found \n", r.RequestURI)
}

func (e *Engine) Run() {
	// user key :get value: func
	//for _, group := range e.routerGroup {
	//	for key, value := range group.handlerFuncMap {
	//		http.HandlerFunc("/"+group.name+key, value)
	//		log.Printf("/" + group.name + key)
	//	}
	//}
	http.Handle("/", e)
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
