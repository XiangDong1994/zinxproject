package main

import (
	"zinx/net"
	"fmt"
	"zinx/ziface"
)

//继承于baseRouter的对象
type PingRouter struct {
	net.BaseRouter
}

//PreHandle方法  ---  用户可以在处理业务之前  自定义一些业务， 实现这个方法

//Handler方法  ---- 用户可以定义一个 业务处理的 核心方法

//PostHandle方法  --- 用户可以在处理业务之后 定义一些业务，实现这个方法

//提供自定义的业务方法
func (this *PingRouter) PreHandle(reqeust ziface.IRequest) {
	fmt.Println("Call Router PreHandler...")
	//给客户端回写一个 数据
	_, err := reqeust.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

func (this *PingRouter) Handle(reqeust ziface.IRequest) {
	fmt.Println("Call Router Handler...")
	//给客户端回写一个 数据
	_, err := reqeust.GetConnection().GetTCPConnection().Write([]byte("ping..ping..ping...\n"))
	if err != nil {
		fmt.Println("call  ping error")
	}
}

func (this *PingRouter) PostHandle(reqeust ziface.IRequest) {
	fmt.Println("Call Router PostHandler...")
	//给客户端回写一个 数据
	_, err := reqeust.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}

func main() {
	//建立一个zinx 对象
	s:=net.NewServer("ZinxV0.3")


	s.AddRouter(&PingRouter{})

	//启动服务
	s.Serve()


	return
}
