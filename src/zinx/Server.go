package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

// ping test customised  router
type PingRouter struct {
	znet.BaseRouter
}

// Ping handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	// first read the client data, and echo 'ping...ping...ping...'
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(0,[]byte("ping...ping...ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

// HelloZinxRouter Handle
type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HellZinxRouter Handle")
	// first read the client data, and echo 'ping...ping...ping...'
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1,[]byte("Hello Zinx Router v0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

// do while conn begin
func DoConnectionBegin(conn ziface.IConnection) {
    fmt.Println("DoConnecionBegin is Called ... ")
	fmt.Println("Set conn Name, Home done!")
	conn.SetProperty("Name","Aceld")
	conn.SetProperty("Home","https://www.lanqiao.cn/courses/1639/")

    err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
    if err != nil {
        fmt.Println(err)
    }
}
// do while conn ends
func DoConnectionLost(conn ziface.IConnection) {
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Conn Property Name = ", name)
	}
	if home, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Conn Property Home = ", home)
	}
    fmt.Println("DoConneciotnLost is Called ... ")
}

func main() {
	// Create a server handler
	s := znet.NewServer()
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	s.AddRouter(0,&PingRouter{})
	s.AddRouter(1,&HelloZinxRouter{})
	s.Serve()
}










