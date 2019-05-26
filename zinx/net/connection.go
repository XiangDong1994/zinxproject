package net

import (
	"net"
	"zinx/ziface"
     "errors"
	"fmt"
	"io"
)

//具体的TCP链接模块
type Connection struct {
	//当前链接的原生套接字
	Conn *net.TCPConn

	//链接ID
	ConnID uint32

	//当前链接状态
	isClosed bool

/*	//当前链接所绑定的业务处理方法
	handleAPI ziface.HandleFunc*/

	//当前链接所绑定的Router
	MsgHandler ziface.IMsgHandler

	//添加一个 Reader和Writer通信的Channel|
	msgChan  chan []byte

	//创建一个Channel  用来Reader通知Writer conn已经关闭，需要退出的消息
	writerExitChan chan bool

}

//初始化链接方法
func NewConnection(conn *net.TCPConn, connID uint32, handler ziface.IMsgHandler) ziface.IConnection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		//handleAPI: callback_api,
		MsgHandler:handler,
		isClosed:  false,
		msgChan: make(chan []byte), //初始化Reader Writer通信的Channel
		writerExitChan:make(chan bool),
	}
	return  c
}
//针对链接读业务的方法
func (c *Connection) StartReader() {
	//从对端读数据
	fmt.Println("Reader go is startin....")
	defer fmt.Println("connID = ", c.ConnID, "Reader is exit, remote addr is = ", c.GetRemoteAddr().String())
	defer c.Stop()



	for {

		/*
	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		//将当前一次性得到的对端客户端请求的数据 封装成一个Request
		req := NewRequest(c, buf, cnt)*/

		//调用用户传递进来的业务 模板 设计模式

		//PreHandle
		//handler
		//postHandle
		/*
		//将数据 传递给我们 定义好的Handle Callback方法
		if err := c.handleAPI(req); err != nil {
			fmt.Println("ConnID", c.ConnID, "Handle is error", err)
			break
		}
		*/

	//	创建拆包封包对象
		dp := NewDataPack()

		//读取客户端消息的头部
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			fmt.Println("read msg head error", err)
			break
		}

		//根据头部 获取数据的长度，进行第二次读取
		msg, err := dp.UnPack(headData) //将msg 头部信息填充满
		if err != nil {
			fmt.Println("unpack error ", err)
			break
		}

		//根据长度 再次读取
		var data []byte
		if msg.GetMsgLen() > 0 {
			//有内容
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				fmt.Println("read msg data error  ", err)
				break
			}

		}
		msg.SetData(data)

		//将读出来的msg 组装一个request
		//将当前一次性得到的对端客户端请求的数据 封装成一个Request
		req := NewRequest(c, msg)

		//调用用户传递进来的业务 模板 设计模式

			go c.MsgHandler.DoMsgHandler(req)

	}
	}


/*
 写消息的Goroutine 专门负责给客户端发送消息
 */
 func (c *Connection)StartWriter(){

	 fmt.Println("[Writer Goroutine isStarted]...")
	 defer fmt.Println("[Writer Goroutine Stop...]")
	 //IO多路复用
	 for {
		 select {
		 case data := <-c.msgChan:
			 //有数据需要写给客户端
			 if _, err := c.Conn.Write(data); err != nil {
				 fmt.Println("Send data error ", err)
				 return
			 }
		 case <-c.writerExitChan:
			 //代表reader已经退出了，writer也要退出
			 return
		 }

	 }
 }





//启动链接
func (c *Connection) Start() {

	fmt.Println("Conn Start（）  ... id = ", c.ConnID)
	//先进行读业务
	go c.StartReader()

	//进行写业务
	go c.StartWriter()
}

//停止链接
func (c *Connection) Stop() {
      fmt.Println("stop the connection id is :",c.ConnID)
      //回收工作
      if c.isClosed==true{
      	return
	  }
      c.isClosed = true

      //
      _ = c.Conn.Close()
      //释放channal资源
	close(c.msgChan)
	close(c.writerExitChan)
}

//获取链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取conn的原生socket套接字
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取远程客户端ip地址
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据给客户端
func (c *Connection) Send(msgId uint32,msgData []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed ..send Msg ")
	}
		//封装成msg
		dp := NewDataPack()

		binaryMsg, err := dp.Pack(NewMsgPackage(msgId, msgData))
		if err != nil {
			fmt.Println("Pack error msg id = ", msgId)
			return err
		}

		//将要发送的打包好的二进制数发送channel 让writer去写
		c.msgChan <- binaryMsg

		return nil


}