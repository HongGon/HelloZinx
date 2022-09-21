package main

import (
	"fmt"
	"io"
	"net"
	"zinx/znet"
)

//  unpack
func main() {
	// create  socket TCP Server
	listener, err := net.Listen("tcp","127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err:", err)
	}

	// Create server goroutine
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server accept err:", err)
		}
		// process server request
		go func(conn net.Conn) {
			// create dp object to pack/unpack
			dp := znet.NewDataPack()
			for {
				// 1 read the head
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("read head error")
					break
				}
				// unpack the headData byte stream into msg
				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("server unpack err:", err)
					return
				}
				if msgHead.GetDataLen() > 0 {
					// read data in msg again
					msg := msgHead.(*znet.Message)
					msg.Data = make([]byte, msg.GetDataLen())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("server unpack data err:", err)
						return
					}
					fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
				}
			}
		}(conn)
	}
}





