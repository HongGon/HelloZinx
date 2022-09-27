package ziface

/*
	Message manage
*/

type IMsgHandle interface {
	// process the msg in non-block mode
	DoMsgHandler(request IRequest)
	// process method to add a router
	AddRouter(msgId uint32, router IRouter)
}






