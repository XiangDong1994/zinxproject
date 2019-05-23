package net

import "zinx/ziface"

type BaseRouter struct {

}

//处理业务之前的方法
func (r *BaseRouter)PreHandle(request ziface.IRequest){
	//将interface的方法全部实现， 目的是 让用户重写这个方法
}
//处理业务的主要方法
func (r *BaseRouter)Handle(request ziface.IRequest){
	//将interface的方法全部实现， 目的是 让用户重写这个方法

}
//处理业务之后的方法
func (r *BaseRouter)PostHandle(request ziface.IRequest)  {
	//将interface的方法全部实现， 目的是 让用户重写这个方法
}
