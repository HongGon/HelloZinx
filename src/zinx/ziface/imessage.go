package ziface

/*
	message interface
*/


type IMessage interface {
	// obtain the length of message
	GetDataLen()	uint32
	// obtain the id of message
	GetMsgID()		uint32
	// obtain the content of message
	GetData()		[]byte
	// set the id of message
	SetMsgId(uint32)
	// set the content of message
	SetData([]byte)
	// set the length of message
	SetDataLen(uint32)
}




