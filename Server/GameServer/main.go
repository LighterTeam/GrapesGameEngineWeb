package main

import (
	"fmt"
	"net"
	. "TSTCP"
)

func main() {
	tcp := new(TSTCP)
	tcp.Create_Server(":9901",
		func() {
			fmt.Println("成功...DB成功! Port: 9901")
		},

		func(conn net.Conn) uint64 {
			fmt.Println("Success!")
			return 0
		},

		func(conn net.Conn, sBuffer string, UUID uint64) {
			fmt.Println("Recv:",sBuffer)

			tcp.SendBuffer([]byte(sBuffer))
			fmt.Println("Send:",sBuffer)
		},

		func(conn net.Conn, UUID uint64) {
			fmt.Println("DB客户端下线, UUID: %v", UUID)
		},
	)
}
