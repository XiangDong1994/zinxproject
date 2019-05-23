package ziface

/*
抽象 IRequest 一次性请求的数据封装
*/
type IRequest interface {
	//得到当前的的请求链接
	GetConnection()  IConnection


	//得到链接数据
	GetData() []byte

	//得到链接长度
	GetDataLen()  int
}