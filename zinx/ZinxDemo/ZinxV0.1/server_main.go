package main

import "zinx/net"

func main() {
	//创建一个zinx server对象
	s:=net.NewServer("zinx v0.1")

	//让servers对象  启动服务
	s.Serve()
}
