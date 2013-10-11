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
			tcp.SendBuffer([]byte("1"))
			tcp.SendBuffer([]byte("12345678901234567890"))
			tcp.SendBuffer([]byte("12"))
			tcp.SendBuffer([]byte("abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1abcdefgvwxyz1"))
			tcp.SendBuffer([]byte("abcdefgvwxyz1"))
			tcp.SendBuffer([]byte("123"))
			tcp.SendBuffer([]byte("abcdefghijklmnop2"))
			tcp.SendBuffer([]byte("1234"))
			tcp.SendBuffer([]byte("abcdefghijklmnopqrstuvwxyz3"))
			tcp.SendBuffer([]byte("12345"))
			tcp.SendBuffer([]byte("abcdefghijklmnvwxyz4"))
			tcp.SendBuffer([]byte("123456"))
			tcp.SendBuffer([]byte("abcdefghijklmn"))
//			tcp.SendBuffer([]byte("123451234512345123451234512345123451234512345"))
//			tcp.SendBuffer([]byte("12345的撒的撒12345123451234512345123451234512345"))
//			tcp.SendBuffer([]byte("12345123451234512345123451234512345123451234512345"))
			tcp.SendBuffer([]byte("12345的撒的撒的撒的撒撒水水水水水水水水水水水水水水水水水水水水水水水水水水水水水水水水"))
		},

		func(conn net.Conn, UUID uint64) {
			fmt.Println("DB客户端下线, UUID: %v", UUID)
		},
	)
}
