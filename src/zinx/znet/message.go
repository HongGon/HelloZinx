package znet

type Message struct {
	// id of message
	Id		uint32
	// length of message
	DataLen uint32
	// content of message
	Data	[]byte
}

// Create a message pack
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:		id,
		DataLen: uint32(len(data)),
		Data:	data,
	}
}

// obtain the length of message
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// obtain the id of msg
func (msg *Message) GetMsgID() uint32 {
	return msg.Id
}

// obtain the content of msg
func (msg *Message) GetData() []byte {
	return msg.Data
}

// set the length of msg
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

// set id of msg
func (msg *Message) SetMsgId(msgId uint32) {
	msg.Id = msgId
}

// set the content of msg
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}







