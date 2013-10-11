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
			fmt.Println("成功...DB成功! Port: 9901")
		},

		func(conn net.Conn) uint64 {
			fmt.Println("Success!")
			AddUUID++
			return AddUUID
		},

		func(conn net.Conn, sBuffer string, UUID uint64) {
			fmt.Println("UUID:",UUID," Recv:",sBuffer)

			SendBuffer(conn, []byte(sBuffer))
			fmt.Println("UUID:",UUID," Send:",sBuffer)
		},

		func(conn net.Conn, UUID uint64) {
			fmt.Println("DB客户端下线, UUID: ", UUID)
		},
	)
}
