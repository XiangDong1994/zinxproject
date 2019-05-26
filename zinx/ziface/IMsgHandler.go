package ziface


type IMsgHandler interface {
	//添加map到几何中
	AddRouter(msgID uint32,router IRouter)
	//调度路由  根据msgID
	DoMsgHandler(request IRequest)
}