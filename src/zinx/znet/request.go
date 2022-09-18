package znet

import (
	"zinx/ziface"
)

type Request struct {
	conn ziface.IConnection // connection built with client
	data []byte
}

// obtain the info of request
func(r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func(r *Request) GetData() []byte {
	return r.data
}


