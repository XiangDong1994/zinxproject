package ziface



type IConnManager interface {

	Add(conn IConnection)


	Remove(connID  uint32)


	Get(connID uint32)(IConnection,error)


	Len() int

	ClearConn()//清空全部链接方法
}
