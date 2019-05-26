/**
* @Author: Aceld(刘丹冰)
* @Date: 2019/5/23 17:26
* @Mail: danbing.at@gmail.com
*/
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//全局配置文件的类
type GlobalObj struct {
	/*
	   server
	 */
	Host string //当前监听的IP
	Port  int   //当前的监听Port
	Name  string  //当前zinxserver的名称


	Version string //当前框架的版本号
	MaxPackageSize uint32 //每次Read一次的最大长度
}

//定义一个全局的对外的配置的对象
var GlobalObject *GlobalObj


//添加一个加载配置文件的方法
func (g *GlobalObj) LoadConfig() {
	data, err := ioutil.ReadFile("conf/zinx.json") //针对main主进行的所在路径的相对路径
	if err != nil {
		//panic(err)
		fmt.Println("load config error")
		return
	}

	//将zinx.json 的数据转换到 GlobalObject中， json一个解析过程
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

//只要import 当前模块 就会执行init方法 加载配置文件
func init() {
	//配置文件的读取操作
	GlobalObject = &GlobalObj{
		//默认值
		Name:"ZinxServerApp",
		Host:"0.0.0.0",
		Port:8999,
		Version:"V0.4",
		MaxPackageSize:512,
	}

	//加载文件
	GlobalObject.LoadConfig()
}