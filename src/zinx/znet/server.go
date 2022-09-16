package znet

import (
	"fmt"
	"net"
	"time"
	"zinx/ziface"
)

// realize the iServer interface, define a class of Server
type Server struct {
	// name of server
	Name string
	// tcp4 or other
	IPVersion string
	// IP address
	IP string
	// Port binded
	Port int
}



// start the networks

func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)
    
    // create a gorountine to handle the Listener business
    go func() {
        // 1 obtain a tcp addr
        addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
        if err != nil {
            fmt.Println("resolve tcp addr err: ", err)
            return
        }
        // 2 monitor the addr
        listenner, err:= net.ListenTCP(s.IPVersion, addr)
        if err != nil {
            fmt.Println("listen", s.IPVersion, "err", err)
            return
        }
        // succeed to monitor
        fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")
        //3 start the network business
        for {
            //3.1 block the server, create the request of connection
            conn, err := listenner.AcceptTCP()
            if err != nil {
                fmt.Println("Accept err ", err)
                continue
            }
            //3.2 TODO Server.Start() set the max connection. if exceed the max connections, then close the new conn
            //3.3 TODO Server.Start() process the request of conn
            // 512 bytes echo service
            go func () {
                // obtain the data from client
                for  {
                    buf := make([]byte, 512)
                    cnt, err := conn.Read(buf)
                    if err != nil {
                        fmt.Println("recv buf err ", err)
                        continue
                    }
                    // echo
                    if _, err := conn.Write(buf[:cnt]); err !=nil {
                        fmt.Println("write back buf err ", err)
                        continue
                    }
                }
            }()
        }
    }()
}

func (s *Server) Stop() {
    fmt.Println("[STOP] Zinx server , name " , s.Name)
    //TODO：  Server.Stop() stop and clean other info
}
func (s *Server) Serve() {
    s.Start()
    //TODO： Server.Serve() 
    //block, otherwise go exit, gorountine of listener exits
    for {
        time.Sleep(10*time.Second)
    }
}

/*
Create a handle of server
*/

func NewServer (name string) ziface.IServer {
	s := &Server {
		Name :name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 7777,
	}
	return s
}


