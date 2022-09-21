package main

import (
	"fmt"
	"net"
	"zinx/znet"
)

func main() {
	// client goroutine
	conn, err := net.Dial("tcp","127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:",err)
	}
	// create a pack obj dp
	dp := znet.NewDataPack()
	// pack a msg1 pack
	msg1 := &znet.Message{
		Id:		0,
		DataLen:5,
		Data:	[]byte{'h','e','l','l','o'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err:",err)
		return
	}
	
	msg2 := &znet.Message{
		Id:		1,
		DataLen:7,
		Data:	[]byte{'w','o','r','l','d','!','!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 err:",err)
		return
	}

	// concat the sendData1 and sendData2
	sendData1 = append(sendData1,sendData2...)
	// write the data to server
	conn.Write(sendData1)
	// block the server
	select {}

}





