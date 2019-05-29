package ZinxV0_5

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
	err := reqeust.GetConnection().Send(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}



func main() {
	//建立一个zinx 对象
	s:=net.NewServer("ZinxV0.5")


	s.AddRouter(&PingRouter{})

	//启动服务
	s.Serve()


	return
}
