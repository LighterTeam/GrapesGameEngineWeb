package main

import (
	"fmt"
	"net"
	. "TSTCP"
)

var AddUUID uint64

func main() {
	tcp := new(TSTCP)
	tcp.Create_Server(":9901",
		func() {
			fmt.Println("TCPServer Success! Port: 9901")
		},

		func(conn net.Conn) uint64 {
			AddUUID++
			fmt.Println("Client connect! UUID:", AddUUID)
			return AddUUID
		},

		func(conn net.Conn, sBuffer string, UUID uint64) {
			fmt.Println("UUID:",UUID," Recv:",sBuffer)

			SendBuffer(conn, []byte(sBuffer))
			fmt.Println("UUID:",UUID," Send:",sBuffer)
		},

		func(conn net.Conn, UUID uint64) {
			fmt.Println("Client close connect! UUID: ", UUID)
		},
	)
}
