package main

import (
	"zinx/znet"
)


//  test func of Server module
func main() {
	// 1 Create a server handler s
	s := znet.NewServer("[zinx v0.1]")
	// 2 Start the service
	s.Serve()
}


