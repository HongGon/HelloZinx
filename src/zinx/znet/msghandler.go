package znet

import (
	"fmt"
	"strconv"
	"zinx/ziface"
)

type MsgHandle struct {
	// save the map attr of processing method according to MsgId
	Apis map[uint32] ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}


func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgID(), " is not FOUND!")
		return 
	}
	// process method
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}


func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	// 1 judge if API process method exists
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api , msgId = " + strconv.Itoa(int(msgId)))
	}

	// 2 add the bind relation of msg and api
	mh.Apis[msgId] = router
	fmt.Println("Add api msgId = ", msgId)
}











