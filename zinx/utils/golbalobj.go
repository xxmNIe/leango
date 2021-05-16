package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"leango/zinx/ziface"
)


//zinx框架的全局参数
type GlobalObj struct {

	/*
	server
	 */
	TopServer ziface.IServer  //当前zinx全局server对象
	Host string
	TcpPort int
	Name string


	/*
	zinx
	 */

	Version string
	MaxConn int
	//connection 中buf的size
	MaxPackageSize uint32    //数据包的最大值

	WorkPoolSize uint32  //当前业务工作worker池数量
	MaxWorkerTaskLen uint32  //允许最大的
}

var GlobalObject *GlobalObj

func(g *GlobalObj)Reload(){

	data,err :=ioutil.ReadFile("/home/xxm/leango/myDemo/V9/conf/zinx.json")

	//解析
	if err !=nil{
		panic(err)
	}
	err = json.Unmarshal(data,&GlobalObject)
	if err!=nil{
		panic(err)
	}

}


//默认舒适化 导包就会执行  且在main之前
func init(){
	fmt.Println("global init()....")
	//如果配置未见没加载，这就是默认的
	GlobalObject = &GlobalObj{
		Name: "ZinxServerApp",
		Version: "V0.4",
		TcpPort: 8999,
		Host: "0.0.0.0",
		MaxConn: 1000,
		MaxPackageSize: 4096,
	}

	GlobalObject.Reload()
}

