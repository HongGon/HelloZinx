package znet

import (
	"zinx/ziface"
)

type BaseRouter struct {}


func (br *BaseRouter) PreHandle(req ziface.IRequest){}
func (br *BaseRouter) Handle(req ziface.IRequest){}
func (br *BaseRouter) PoseHandle(req ziface.IRequest){}






