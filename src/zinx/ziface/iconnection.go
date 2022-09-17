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
	// // Server on serve
	// Serve()
}


// define a interface of connection business
type HandFunc func(*net.TCPConn, []byte, int) error
