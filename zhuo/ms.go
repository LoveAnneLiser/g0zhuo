package zhuo

import (
	"fmt"
	"log"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

// 分组路由
type routerGroup struct {
	name             string
	handleFuncMap    map[string]HandleFunc
	handlerMethodMap map[string][]string
}

func (r *routerGroup) Add(name string, handleFunc HandleFunc) {
	r.handleFuncMap[name] = handleFunc
}

func (r *routerGroup) Any(name string, handleFunc HandleFunc) {
	r.handleFuncMap[name] = handleFunc
	r.handlerMethodMap["ANY"] = append(r.handlerMethodMap["ANY"], name)
}

func (r *routerGroup) Get(name string, handleFunc HandleFunc) {
	r.handleFuncMap[name] = handleFunc
	r.handlerMethodMap[http.MethodGet] = append(r.handlerMethodMap[http.MethodGet], name)
}

func (r *routerGroup) Post(name string, handleFunc HandleFunc) {
	r.handleFuncMap[name] = handleFunc
	r.handlerMethodMap[http.MethodPost] = append(r.handlerMethodMap[http.MethodPost], name)
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
		handleFuncMap:    make(map[string]HandleFunc),
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
				routes, ok := group.handlerMethodMap["ANY"]
				if ok {
					for _, routerName := range routes {
						if routerName == name {
							methodHandle(w, r)
							return
						}
					}
				}
				// 根据method进行匹配
				routes, ok = group.handlerMethodMap[method]
				if ok { // 代表路径一样 但是请求方式没有
					for _, routerName := range routes {
						if routerName == name {

							methodHandle(w, r)
							return
						}
					}
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
