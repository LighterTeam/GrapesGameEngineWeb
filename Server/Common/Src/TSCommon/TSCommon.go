package TSCommon

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Println("TSCommon init");
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var GateWayServer_IP string = "127.0.0.1"				// 网关服务器IP
var GateWayServer_Port int = 9001						// 网关服务器Port

var DBServer_IP string = "127.0.0.1"						// DB前端IP
var DBServer_Port int = 9002								// DB前端Port

var HTTPClientServer_Port int = 8080

const Max_Pool_GameServer int = 100						// 游戏服上限




