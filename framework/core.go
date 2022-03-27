package framework

import "net/http"

// 框架核心结构
type Core struct {
}

// 初始化框架核心结构
func NewCore() *Core {
	return &Core{}
}

// 框架核心结构时间handler接口
func (c *Core) ServerHttp(response http.ResponseWriter, request *http.Request) {
	// TODO
}
