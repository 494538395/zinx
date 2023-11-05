package net

import (
	"fmt"
	"github.com/494538395/zinx/iface"
)

// 实现 Router 接口时，先嵌入这个 BaseRouter 基类，然后根据需求再对这个基类进行重写就好
type BaseRouter struct {
}

// 这是所有 BaseRouter 的方法都是空实现，是因为 BaseRouter 起到了一个 JAVA 中抽象类到作用
// 比如：有些 Router 不需要 PreHandle 和 PostHandle ，那它就只用拥有 BaseRouter 属性，然后重写 BaseRouter 到 Handle 即可，不需要再
// 手动实现 PreHandle 和 PostHandle （因为 BaseRouter 都已经实现过了）

// PreHandle 处理业务前的 hook
func (r *BaseRouter) PreHandle(req iface.IRequest) {
	fmt.Println("BaseRouter 的默认 PreHandle 实现")
}

// Handle 处理业务的方法
func (r *BaseRouter) Handle(req iface.IRequest) {
	fmt.Println("BaseRouter 的默认 Handle 实现")
}

// PostHandle 处理业务之后的 hook
func (r *BaseRouter) PostHandle(req iface.IRequest) {
	fmt.Println("BaseRouter 的默认 PostHandle 实现")
}
