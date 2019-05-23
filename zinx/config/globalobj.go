package config
import   (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

//全局配置文件的类

type Globalobj struct {
/*
server
*/
Host string //当前监听的IP
Post int//监听的post
Name string//监听Name

Version string //当前版本号
MaxPackageSize  uint32 //每次read的最大长度
}

//创建一个全局对象
var Globalobject *Globalobj

//添加加载文件配置的方法
func(g *Globalobj)LoadConfig(){

	data,err:=ioutil.ReadFile(Globalobject.Name)
	if err!=nil{
		fmt.Println("config ioutil err is:",err)
		return
	}

	//将zinx.json 的数据转换到 GlobalObject中， json一个解析过程
	err = json.Unmarshal(data,&Globalobject)
	if err != nil {
		panic(err)
	}

}
//只要import当前模块 就会执行init 方法  加载配置文件
func init(){
	Globalobject = &Globalobj{
		//默认值
		Name:"ZinxServerApp",
		Host:"0.0.0.0",
		Post:8999,
		Version:"V0.4",
		MaxPackageSize:512,
	}

	Globalobject.LoadConfig()
}