package ziface


import "net"


// define the connection interface
type IConnection interface{
	// Start the server
	Start()
	// Stop the server
	Stop()
	// obtain raw socket TCPConn from current connection
	GetTCPConnection() *net.TCPConn
	// obtain the current conn id
	GetConnID() uint32
	// obtain addr info of remote client
	RemoteAddr() net.Addr
	//  Send Msg
	SendMsg(msgId uint32, data []byte) error
	// send msg data to remote TCP client (buffer)
	SendBuffMsg(msgId uint32, data []byte) error
	// Set the property of conn
	SetProperty(key string, value interface{})
	// get the property of conn
	GetProperty(key string)(interface{}, error)
	//  remove the property of conn
	RemoveProperty(key string)

	// // Server on serve
	// Serve()
}


// define a interface of connection business
type HandFunc func(*net.TCPConn, []byte, int) error
