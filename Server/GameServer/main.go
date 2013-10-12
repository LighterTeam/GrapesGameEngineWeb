package main

import (
	"fmt"
	"net"
	. "TSTCP"
	"TSLog"
	"TSEngine"
)

var AddUUID uint64

func main() {
	//Log
	TSLog.InitLogger("ServerLog")
	TSLog.SetLevel(TSLog.TraceLevel)
	fmt.Println("ServerLog init Success!")

	tcp := new(TSTCP)
	tcp.Create_Server(":9901",
		func() {
			fmt.Println("TCPServer Success! Port: 9901")
		},

		func(conn net.Conn) uint64 {
			AddUUID++
			fmt.Println("Client connect! UUID:", AddUUID)
			TSEngine.RegistClient(AddUUID, conn)
			return AddUUID
		},

		func(conn net.Conn, sBuffer []byte, UUID uint64) {
			fmt.Println("UUID:",UUID," Recv:",string(sBuffer))
			TSEngine.ProcessMsg(sBuffer, UUID)
		},

		func(conn net.Conn, UUID uint64) {
			fmt.Println("Client close connect! UUID: ", UUID)
			TSEngine.UnregistClient(UUID)
		},
	)
}
