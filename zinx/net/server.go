/*
servser模块实现层
*/


package net

import (
	"zinx/ziface"
	"fmt"
	"net"
)

type Server struct {
	//服务器IP
	IPVersion string
	IP string
	//服务器接口
	Port int
	// 服务器名称
	Name string

	Router ziface.IRouter
}
//定义回显业务
func CallBackBusi(request ziface.IRequest) error {
	//回显业务
	fmt.Println("【conn Handle】 CallBack..")
	c := request.GetConnection().GetTCPConnection()
	buf := request.GetData()
	cnt := request.GetDataLen()
	if _, err := c.Write(buf[:cnt]);err !=nil {
		fmt.Println("write back err ", err)
		return err
	}

	return nil
}


//初始化的new方法

func NewServer(name string)ziface.IServer{
	s := &Server{
		Name:name,
		IPVersion:"tcp4",
		IP:"0.0.0.0",
		Port:8999,
		Router:nil,

	}
	return s
}



//启动服务器
//创建原生的socket
func (s *Server) Start() {
	fmt.Printf("[start] Server Linstenner at IP :%s, Port :%d, is starting..\n", s.IP, s.Port)

	//1 创建套接字  ：得到一个TCP的addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error:", err)
		return
	}
	//2 监听服务器地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen ", s.IPVersion, " err , ", err)
		return
	}

	//生成id的累加器
	var cid uint32
	cid = 0

	//3 阻塞等待客户端发送请求，
	go func() {
		for {
			//阻塞等待客户端请求,
			conn, err := listenner.AcceptTCP()//只是针对TCP协议
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//创建一个Connection对象
			dealConn := NewConnection(conn, cid, s.Router)
			cid++


			//此时conn就已经和对端客户端连接
			go dealConn.Start()
		}
		/*	go func() {
		//4客户端有数据请求，处理客户端业务	（读写)
		for {
			buf := make([]byte, 512)
			cnt, err := conn.Read(buf)
			if err != nil {
				fmt.Println("recv buf.err", err)
				break

			}
			fmt.Printf("recv client buf is %s,cnt = %d\n", buf, cnt)
			//回显功能（业务）"lien
			_, err = conn.Write(buf[:cnt])
			if err != nil {
				fmt.Println("write err is:", err)
				continue
			}
		}
	}()*/


	}()

}



//停止服务器
func(s *Server)Stop(){

	//todo 将一些资源进行回收
}
//运行服务器
func(s *Server)Serve(){
//启动监听功能
s.Start()//并不希望他永久阻塞
//todo 做一些其他扩展
 //阻塞//告诉cpu不再需要处理的，节省cpu资源
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}



