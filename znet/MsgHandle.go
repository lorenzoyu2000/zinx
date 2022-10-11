package znet

import (
	"fmt"
	"github.com/lorenzoyu2000/zinx/utils"
	"github.com/lorenzoyu2000/zinx/ziface"
)

type MsgHandle struct {
	// MsgID --> router方法
	Apis map[uint32]ziface.IRouter
	// 工作的消息队列
	TaskQueue []chan ziface.IRequest
	// 工作Goroutine的数量
	WorkPoolSize uint32
}

func (mh *MsgHandle) DoMsgRouter(request ziface.IRequest) {
	router, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("MsgID: ", request.GetMsgId(), " is not found")
		return
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	_, ok := mh.Apis[msgId]
	if ok {
		fmt.Println("MsgID: ", msgId, " is already exist")
		return
	}
	mh.Apis[msgId] = router
	fmt.Println("MsgID: ", msgId, " add succeed")
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:         make(map[uint32]ziface.IRouter),
		WorkPoolSize: utils.GlobalObject.WorkPoolSize,
		TaskQueue:    make([]chan ziface.IRequest, utils.GlobalObject.WorkPoolSize),
	}
}

func (mh *MsgHandle) StartWorkPool() {
	// 启动WorkPoolSize个Goroutine来接收消息
	for i := 0; i < int(mh.WorkPoolSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxTaskWorkLen)
		go mh.startOneWorker(i)
	}
}

func (mh *MsgHandle) startOneWorker(workID int) {
	fmt.Println("WorkID =", workID, " is started...")

	for {
		select {
		case request := <-mh.TaskQueue[workID]:
			mh.DoMsgRouter(request)
		}
	}
}

func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// TODO 优化负载均衡算法
	workID := request.GetConnection().GetConnID() % mh.WorkPoolSize
	fmt.Println("ConnID", request.GetConnection().GetConnID(), " MsgID", request.GetMsgId(), " to WorkID", workID)

	mh.TaskQueue[workID] <- request
}
