package main
import (
    "fmt"
    "io"
    "net"
    "time"
    "zinx/znet"
)
/*
	Simulate the client

*/
func main() {
    fmt.Println("Client Test ... start")
    // launch the request after 3s
    time.Sleep(3 * time.Second)
    conn,err := net.Dial("tcp", "127.0.0.1:7777")
    if err != nil {
        fmt.Println("client start err, exit!")
        return
    }
    for {
        // send msg
        dp := znet.NewDataPack()
        msg, _ := dp.Pack(znet.NewMsgPackage(1,[]byte("Zinx V0.6 Client1 Test Message")))
        _, err := conn.Write(msg)
        if err !=nil {
            fmt.Println("write error err ", err)
            return
        }
        //  read head
        headData := make([]byte, dp.GetHeadLen())
        _, err = io.ReadFull(conn, headData)
        if err != nil {
            fmt.Println("read head error")
            break
        }
        // unpack
        msgHead, err := dp.Unpack(headData)
        if err != nil {
            fmt.Println("server unpack err:", err)
            return
        }
        if msgHead.GetDataLen() > 0 {
            //msg
            msg := msgHead.(*znet.Message)
            msg.Data = make([]byte, msg.GetDataLen())
            //read byte stream
            _, err := io.ReadFull(conn, msg.Data)
            if err != nil {
                fmt.Println("server unpack data err:", err)
                return
            }
            fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
        }
        time.Sleep(1*time.Second)
    }
}