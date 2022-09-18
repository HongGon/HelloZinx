package ziface

/*
	Router interface
*/

type IRouter interface {
	PreHandle(request IRequest)		// the handler method before processing the conn business 
	Handle(request IRequest)		// the method to process the conn business
	PostHandle(request IRequest)	// the handler method after processing the conn business
}





