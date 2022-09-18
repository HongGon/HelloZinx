package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	// Current socket TCP
	Conn *net.TCPConn
	// Current ID (SessionID), globally unique
	ConnID uint32
	// close state
	isClosed bool
	
	// // api of process method for this conn
	// handleAPI ziface.HandFunc

	// process method of router
	Router ziface.IRouter
	// notice that this conn has exited
	ExitBuffChann chan bool
}

// method to create a conn
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection{
	c := &Connection{
		Conn:		conn,
		ConnID:		connID,
		isClosed:	false,
		// handleAPI:	callback_api,
		Router: router,
		ExitBuffChann: make(chan bool, 1),
	}
	return c
}


// the Goroutine to process conn data
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(),"conn reader exit!")
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChann <- true
			continue
		}
		// if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		// 	fmt.Println("connID ", c.ConnID, " handle is error")
		// 	c.ExitBuffChann <- true
		// 	return
		// }

		req := Request{
			conn: c,
			data: buf,
		}

		go func (request ziface.IRequest) {
			// router method to register
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

//  Start conn, let current conn start working
func (c *Connection) Start() {
	// start to process the request after reading client data
	go c.StartReader()
	for {
		select {
		case <- c.ExitBuffChann:
			// get the msg of exit, dont block
			return
		}
	}
}


// Stop conn
func (c *Connection) Stop() {
	// 1. if current conn has closed
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	// TODO Connection Stop() if user register the business of closing the echo, then display the registion
	// close socket conn
	c.Conn.Close()
	// tells the business that read data from buffer queue that this conn has closed
	c.ExitBuffChann <- true
	close(c.ExitBuffChann)
}

// obtain the raw socket from current conn, TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// obtain the id of current conn
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// obtain the addr of remote client
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}













