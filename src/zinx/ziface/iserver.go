package ziface


// define the server interface
type IServer interface{
	// Start the server
	Start()
	// Stop the server
	Stop()
	// Server on serve
	Serve()
	// Router function
	AddRouter(msgId uint32, router IRouter)
}
