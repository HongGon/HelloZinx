package main
import (
    "fmt"
    "zinx/ziface"
    "zinx/znet"
)
//ping test define router
type PingRouter struct {
    znet.BaseRouter
}
//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
    fmt.Println("Call PingRouter Handle")
    // read client data
    fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
    // write data
    err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
    if err != nil {
        fmt.Println(err)
    }
}
func main() {
    // create a handler
    s := znet.NewServer("[zinx v0.5]")
    // config router
    s.AddRouter(&PingRouter{})
    // start the service
    s.Serve()
}