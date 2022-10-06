package znet

import (
	"fmt"
	"github.com/lorenzoyu2000/zinx/ziface"
)

type MsgHandle struct {
	// MsgID --> router方法
	Apis map[uint32]ziface.IRouter
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
		Apis: make(map[uint32]ziface.IRouter),
	}
}
