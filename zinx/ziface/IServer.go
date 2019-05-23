/*
Server模块的抽象层
*/
package ziface

type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Serve()
	//增加添加路由功能   暴露给开发者
	AddRouter(router IRouter)


}

