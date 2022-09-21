package ziface
/*
	IRequest interface
*/

type IRequest interface{
	GetConnection() IConnection
	GetData() []byte
	GetMsgID() uint32
}



