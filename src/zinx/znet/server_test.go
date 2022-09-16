package znet
import (
    "fmt"
    "net"
    "testing"
    "time"
)
/*
    simulate the client
*/
func ClientTest() {
    fmt.Println("Client Test ... start")
    // request test after 3s 
    time.Sleep(3 * time.Second)
    conn,err := net.Dial("tcp", "127.0.0.1:7777")
    if err != nil {
        fmt.Println("client start err, exit!")
        return
    }
    for {
        _, err := conn.Write([]byte("hello ZINX"))
        if err !=nil {
            fmt.Println("write error err ", err)
            return
        }
        buf :=make([]byte, 512)
        cnt, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read buf error ")
            return
        }
        fmt.Printf(" server call back : %s, cnt = %d\n", buf,  cnt)
        time.Sleep(1*time.Second)
}
}


// test function of Server module
func TestServer(t *testing.T) {
    /*
        Server test
    */
    //1 Create a handler of server
    s := NewServer("[zinx V0.1]")
    /*
        Server test
    */
    go ClientTest()
    //2 Serve the server
    s.Serve()
}