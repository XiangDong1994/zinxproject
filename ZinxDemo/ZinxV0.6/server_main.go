/**
 zinx v0.1 应用
*/
package main

import (
	"fmt"
	"zinx/net"
	"zinx/ziface"
)

//PreHandle方法  ---  用户可以在处理业务之前  自定义一些业务， 实现这个方法
//Handler方法  ---- 用户可以定义一个 业务处理的 核心方法
//PostHandle方法  --- 用户可以在处理业务之后 定义一些业务，实现这个方法
type PingRouter struct {
	net.BaseRouter
}

//200 ---> pingpingping
//201 ---> hello zinx..


func (this *PingRouter) Handle(reqeust ziface.IRequest) {
	fmt.Println("Call Router Handler...")
	//给客户端回写一个 数据
	err := reqeust.GetConnection().Send(200, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	net.BaseRouter
}


func (this *HelloRouter) Handle(reqeust ziface.IRequest) {
	fmt.Println("Call Router Handler...")
	//给客户端回写一个 数据
	err := reqeust.GetConnection().Send(201, []byte("hello Zinx!!!"))
	if err != nil {
		fmt.Println(err)
	}
}




func main() {
	//创建一个zinx server对象
	s := net.NewServer("zinx v0.5")

	//注册一些自定义的业务  客户端发送不同的消息，我们服务器应该能够处理不同的业务
	s.AddRouter(1, &PingRouter{}) //真正处理核心业务的
	s.AddRouter(2, &HelloRouter{})

	//让server对象 启动服务
	s.Serve()

	return
}
