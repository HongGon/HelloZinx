package znet

import (
	"fmt"
	"strconv"
	"zinx/ziface"
	"zinx/utils"
)

type MsgHandle struct {
	// save the map attr of processing method according to MsgId
	Apis map[uint32] ziface.IRouter
	// the num of worker
	WorkerPoolSize uint32
	// Task Queue
	TaskQueue []chan ziface.IRequest
}


func (mh *MsgHandle) StartOneWorker(workID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID = ", workID, " is started.")
	// wait for msg in queue
	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}


// launch the worker pool
func (mh *MsgHandle) StartWorkerPool() {
	for i:=0; i<int(mh.WorkerPoolSize);i++ {
		// a worker has been launched
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// launch current worker, block the task queue
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// send msg to task queue
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// allocate the worker according to ConnID
	workerID := request.GetConnection().GetConnID()%mh.WorkerPoolSize
	fmt.Println("Add ConnID=",request.GetConnection().GetConnID(),"request msgID=", request.GetMsgID(),"to workerID=",workerID)
	// sned msg to task queue
	mh.TaskQueue[workerID] <- request
}


func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		// a worker a queue
		TaskQueue: make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
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











