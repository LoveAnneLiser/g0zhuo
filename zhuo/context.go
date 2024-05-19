package zhuo

import "net/http"

// Context 上下文
type Context struct {
	W http.ResponseWriter
	R *http.Request
}
