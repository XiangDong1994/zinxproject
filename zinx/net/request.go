package net

import (
	"zinx/ziface"

)

type Request struct {
	//链接信息
	conn ziface.IConnection

/*	//数据内容
	data []byte
	//数据长度
	len int*/

	msg ziface.IMessage

}
func NewRequest(conn ziface.IConnection,msg  ziface.IMessage)ziface.IRequest{

	req:=&Request{
		conn:conn,
		//data:data,
		//len:len,
		msg:msg,

	}
	return req
}
//得到当前的请求的链接
func(r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

/*//得到链接的数据
func(r *Request) GetData() []byte {
	return r.data
}

//得到链接的长度
func(r *Request)  GetDataLen() int {
	return r.len*/
	func (r *Request)GetMsg()ziface.IMessage{
		return  r.msg
}


//