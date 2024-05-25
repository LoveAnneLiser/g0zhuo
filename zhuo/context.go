package zhuo

import "net/http"

// Context 上下文
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) HTML(status int, html string) error {
	// 设置状态是200 默认不设置的话 如果调用了write方法 实际上默认返回状态200
	// 返回这个html页面
	c.W.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.W.WriteHeader(status)
	_, err := c.W.Write([]byte(html))
	return err
}
