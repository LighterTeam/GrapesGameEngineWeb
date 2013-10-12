package TSEngine

import (
	"sync"
	"net"
	"TSJson"
	. "TSTCP"
	"TSLog"
	"fmt"
)

const (
	SendBroad = iota   //广播
	SendSingle			//单发
	SendBroadNotSelf	//广播没有自己
)

var G_PoolClient map[uint64]*Client = make(map[uint64]*Client)
var G_GameMutex sync.Mutex

func RegistClient(uuid uint64, conn net.Conn){
	G_GameMutex.Lock()
	G_PoolClient[uuid] = &Client{uuid, Sprite{0,0,0}, conn}
	G_GameMutex.Unlock()
	packet := fmt.Sprintf(`{"UUID":%v,"OPCODE":0}`,uuid)
	SendBuffer(conn,[]byte(packet))
}

func UnregistClient(uuid uint64){
	G_GameMutex.Lock()
	delete(G_PoolClient, uuid)
	G_GameMutex.Unlock()
}

func ProcessMsg(data []byte, uuid uint64) {
	var mp MessagePacket
	err := TSJson.Unmarshal(data, &mp)
	if err != nil {
		TSLog.Debug("数据非对应MessagePacket Error! UUID:%v Data:%v", uuid, string(data))
		return
	}

	TSLog.Debug("ProcessMsg: %v",mp)

	G_GameMutex.Lock()
	if mp.T == SendBroad {
		//广播
		for _,v := range(G_PoolClient) {
			SendBuffer(v.Conn, []byte(mp.D))
		}

	} else if mp.T == SendSingle {
		//单发
		SendBuffer(G_PoolClient[mp.S].Conn, []byte(mp.D))
	} else if mp.T == SendBroadNotSelf {
		//广播但是没有自己
		for _,v := range(G_PoolClient) {
			if uuid == v.Uuid {
				continue
			}
			SendBuffer(v.Conn, []byte(mp.D))
		}
	}
	G_GameMutex.Unlock()
}



