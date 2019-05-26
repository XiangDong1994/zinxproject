package net

import (
	"zinx/ziface"
	"fmt"
)

type MsgHandler struct {
	Apis map[uint32]  ziface.IRouter
}

//初始化方法
func NewMsgHandler() ziface.IMsgHandler {
	//给map开辟头空间
	return &MsgHandler{
		Apis:make(map[uint32]ziface.IRouter),
	}
}

//添加路由到map集合中
func (mh *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	//1 判断新添加的msgID key是否已经存在
	if _, ok := mh.Apis[msgID]; ok {
		//msgId已经注册
		fmt.Println("repeat Api msgID = ", msgID)
		return
	}
	//2 添加msgID 和 router的对应关系
	mh.Apis[msgID] = router
	fmt.Println("Apd api MsgID = ", msgID, " succ!")
}

//调度路由， 根据MsgID
func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	// 1 从Request 取到MsgiD
	router, ok := mh.Apis[request.GetMsg().GetMsgId()]
	if !ok {
		fmt.Println("api MsgID = ", request.GetMsg().GetMsgId(), " Not Found! Need Add！")
		return
	}
	//2 根据msgID  找到对应的router 进行调用
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}