package znet

import  (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/utils"
	"zinx/ziface"
)

// instance of pack and unpack
type DataPack struct {}

// initialize method of pack and unpack
func NewDataPack() *DataPack {
	return &DataPack{}
}

// method of obtaining the length
func (dp *DataPack) GetHeadLen() uint32 {
	// Id uint32(4 bytes) + DataLen uint32(4 bytes)
	return 8
}

// pack method
func(dp *DataPack) Pack(msg ziface.IMessage)([]byte, error) {
	// create a buffer to storage the bytes
	dataBuff := bytes.NewBuffer([]byte{})
	// dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	//  write msgId
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	// write data
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil,err
	}
	return dataBuff.Bytes(),nil
}

// unpack method
func(dp *DataPack) Unpack(binaryData []byte)(ziface.IMessage, error) {
	//  create a ioReader
	dataBuff := bytes.NewReader(binaryData)
	// only unpack the info of head to obtain the dataLen and msgID
	msg := &Message{}
	// read dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	// read msgId
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	// judge if length suppress the max
	if (utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize) {
		return nil, errors.New("Too large msg data received")
	}
	// unpack the head data
	return msg,nil
}













