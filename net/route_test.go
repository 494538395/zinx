package net

import (
	"fmt"
	"github.com/494538395/zinx/iface"
	"testing"
)

type MyRouter struct {
	BaseRouter
}

func (mr *MyRouter) Handle(req iface.IRequest) {
	fmt.Println("MyRouter 实现的 Handle")
}

func TestRouter(t *testing.T) {

	m := &MyRouter{}
	m.PreHandle(nil)
	m.Handle(nil)
	m.PostHandle(nil)

}
