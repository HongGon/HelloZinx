package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

/*

	Storage global parameters about Zinx

*/

type GlobalObj struct {
	// global Server of current Zinx
	TcpServer ziface.IServer
	// IP of current server
	Host string
	// port of current server
	TcpPort int
	// name
	Name string
	// version
	Version string
	// max of data
	MaxPacketSize uint32
	// allowed max conn number of current server
	MaxConn int
	// the num of worker
	WorkerPoolSize uint32
	// max of task
	MaxWorkerTaskLen uint32
	// config file path
	ConfFilePath string
}

/*
	define a global object
*/
var GlobalObject *GlobalObj

//  read the config file
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	// parse the json data into struct
	err = json.Unmarshal(data,&GlobalObject)
	if err != nil {
		panic(err)
	}
}


/*
	init method
*/

func init() {
	// initialize the variable GlobalObject
	GlobalObject = &GlobalObj{
		Name:			"ZinxServerApp",
		Version: 		"V0.6",
		TcpPort: 		7777,
		Host:			"0.0.0.0",
		MaxConn:		12000,
		MaxPacketSize: 	4096,
		ConfFilePath: "conf/zinx.json",
		WorkerPoolSize: 10,
		MaxWorkerTaskLen: 1024,
	}
	// load some user config from json
	GlobalObject.Reload()
}













