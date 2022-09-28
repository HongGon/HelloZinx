package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/ziface"
	"zinx/utils"
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
	// Router ziface.IRouter

	// msg manage module
	MsgHandler ziface.IMsgHandle

	// notice that this conn has exited
	ExitBuffChann chan bool

	// non-buffer channel
	msgChan		chan []byte
}

// method to create a conn
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection{
	c := &Connection{
		Conn:		conn,
		ConnID:		connID,
		isClosed:	false,
		// handleAPI:	callback_api,
		// Router: router,
		MsgHandler: msgHandler,
		ExitBuffChann: make(chan bool, 1),
		msgChan: make(chan []byte),
	}
	return c
}


/*
	Goroutine of write msg
*/

func  (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running!]")
	defer fmt.Println(c.RemoteAddr().String(),"[conn Writer exit!]")
	for {
		select{
			case data := <-c.msgChan:
				// data need to be send
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Data error:, ", err , " Conn Writer ")
					return
				}
			case <- c.ExitBuffChann:
				// conn has closed
				return
		}
	}
}





// the Goroutine to process conn data
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(),"conn reader exit!")
	defer c.Stop()
	for {
		// // buf := make([]byte, 512)
		// buf := make([]byte, utils.GlobalObject.MaxPacketSize)
		// _, err := c.Conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("recv buf err ", err)
		// 	c.ExitBuffChann <- true
		// 	continue
		// }
		// if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		// 	fmt.Println("connID ", c.ConnID, " handle is error")
		// 	c.ExitBuffChann <- true
		// 	return
		// }

		// create the obj to pack and unpack
		dp := NewDataPack()
		// read Msg head from client
		headData := make([]byte, dp.GetHeadLen())

		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error ", err)
			// c.ExitBuffChann <- true
			// continue
			break
		}

		// unpack
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error ", err)
			// c.ExitBuffChann <- true
			// continue
			break
		}
		
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err )
				// c.ExitBuffChann <- true
				continue
			}
		}
		msg.SetData(data)
		

		req := Request{
			conn: c,
			msg: msg,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			// send msg to task queue
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandler(&req)
		}

		
		// go func (request ziface.IRequest) {
		// 	// router method to register
		// 	c.Router.PreHandle(request)
		// 	c.Router.Handle(request)
		// 	c.Router.PostHandle(request)
		// }(&req)
	}
}

//  Start conn, let current conn start working
func (c *Connection) Start() {
	// start to process the request after reading client data
	go c.StartReader()
	go c.StartWriter()
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

//  send msg
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}
	// pack data and send it
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}
	// client
	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Println("Write msg id ", msgId, " error ")
		c.ExitBuffChann <- true
		return errors.New("conn Write error")
	}
	return nil
}











