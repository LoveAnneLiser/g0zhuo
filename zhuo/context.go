package zhuo

import (
	"html/template"
	"net/http"
)

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

func (c *Context) HTMLTemplate(name string, data any, filenames ...string) error {
	// 设置状态是200 默认不设置的话 如果调用了write方法 实际上默认返回状态200
	// 返回这个html页面
	c.W.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.New(name)
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		return err
	}
	err = t.Execute(c.W, data)
	return err
}

func (c *Context) HTMLTemplateGlob(name string, data any, pattern string) error {
	// 设置状态是200 默认不设置的话 如果调用了write方法 实际上默认返回状态200
	// 返回这个html页面
	c.W.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.New(name)
	t, err := template.ParseGlob(pattern)
	if err != nil {
		return err
	}
	err = t.ExecuteTemplate(c.W, name, data)
	return err
}
