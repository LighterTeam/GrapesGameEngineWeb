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
var G_RootClient *Client

func RegistClient(uuid uint64, conn net.Conn){
	G_GameMutex.Lock()
	G_PoolClient[uuid] = &Client{uuid, Sprite{0,0,0}, conn}

	if G_RootClient == nil {
		G_RootClient = G_PoolClient[uuid]
	}

	// 给自己注册
	packet := fmt.Sprintf(`{"UUID":%v,"ROOTUUID":%v,"OPCODE":0}`,uuid,G_RootClient.Uuid)
	SendBuffer(conn,[]byte(packet))

	// 给新开启的玩家 创建其他玩家
	for _,v := range(G_PoolClient) {
		if v.Uuid == uuid {
			continue
		}
		packet = fmt.Sprintf(`{"UUID":%v,"OPCODE":1}`,v.Uuid)
		SendBuffer(conn,[]byte(packet))
	}

	// 其他玩家 创建新上线的玩家
	packet = fmt.Sprintf(`{"UUID":%v,"OPCODE":1}`,uuid)
	for _,v := range(G_PoolClient) {
		if v.Uuid != uuid {
			SendBuffer(v.Conn,[]byte(packet))
		}
	}
	G_GameMutex.Unlock()
}

func UnregistClient(uuid uint64){
	G_GameMutex.Lock()
	delete(G_PoolClient, uuid)

	if uuid == G_RootClient.Uuid {
		G_RootClient = nil
		for _,v := range(G_PoolClient) {
			G_RootClient = v
			break
		}

		for _,v := range(G_PoolClient) {
			packet := fmt.Sprintf(`{"ROOTUUID":%v,"OPCODE":4}`,G_RootClient.Uuid)
			SendBuffer(v.Conn,[]byte(packet))
		}
	}

	packet := fmt.Sprintf(`{"UUID":%v,"OPCODE":3}`,uuid)
	for _,v := range(G_PoolClient) {
		SendBuffer(v.Conn,[]byte(packet))
	}
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



