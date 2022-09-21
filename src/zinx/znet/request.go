package znet

import (
	"zinx/ziface"
)

type Request struct {
	conn ziface.IConnection // connection built with client
	msg  ziface.IMessage	// data request by client
	// data []byte
}

// obtain the info of request
func(r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// obtaint the data 
func(r *Request) GetData() []byte {
	return r.msg.GetData()
}


//  obtain the id
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}



