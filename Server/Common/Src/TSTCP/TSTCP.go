package TSTCP

import (
	"bytes"
	"encoding/binary"
	//"fmt"
	"net"
)

type ConnectNew func (conn net.Conn) uint64
type ReceiveBuffer func (conn net.Conn, sBuffer string, UUID uint64)
type ServerInit func ()
type ClientInit func (conn net.Conn)
type ConnectClose func (conn net.Conn, UUID uint64)

type TSTCP struct {
	conn net.Conn
}

func SendBuffer(conn net.Conn, sBuffer []byte) {
	if conn == nil {
		return
	}
	buf := make([]byte, len(sBuffer) + 4)
	binary.BigEndian.PutUint32(buf[:4], uint32(len(sBuffer)))
	copy(buf[4:], sBuffer)
	conn.Write(buf)
}

func (this *TSTCP) SendBuffer(sBuffer []byte) {
	SendBuffer(this.conn, sBuffer)
}

func (this *TSTCP) Create_Server(webpath string, init ServerInit, funCN ConnectNew, funRB ReceiveBuffer, funCC ConnectClose) {
	listener, err := net.Listen("tcp", webpath)
	if err != nil {
		return
	}

	init()

	for {
		this.conn, err = listener.Accept()
		if err != nil {
			return
		}

		UUID := funCN(this.conn)
//		this.SendBuffer([]byte("_RegistUUID," + TSUtil.ToString(int(UUID))))

		go tcpHandler(this.conn, funRB, funCC, UUID)
	}
}

func (this *TSTCP) Create_Client(webpath string, init ClientInit, funRB ReceiveBuffer, funCC ConnectClose) {
	conn, err := net.Dial("tcp", webpath)
	this.conn = conn
	if err != nil {
		return
	}

	init(this.conn)

	go tcpHandler(this.conn, funRB, funCC, 0)
}

func tcpHandler(conn net.Conn, funRB ReceiveBuffer, funCC ConnectClose, UUID uint64) {
	cache := make([]byte, 1024)
	buf := bytes.NewBuffer(make([]byte,0, 1024))

	var contentLen uint32

	for {
		// 读取网络数据
		size, err := conn.Read(cache)
		if err != nil {
			funCC(conn, UUID)
			//fmt.Printf("Read error, %v\n", err.Error())
			return
		}

		// 写入缓冲区
		buf.Write(cache[:size])

		for {
			// 本次缓冲区数据包正好读完，重置内容长度
			if buf.Len() == 0 {
				contentLen = 0
				break
			}

			// 读取一个新的数据包
			if contentLen == 0 {
				// 判断缓冲区剩余数据是否足够读取一个包长
				if buf.Len() < 4 {
					break
				}
				len := make([]byte, 4)
				_, err = buf.Read(len)
				contentLen = binary.BigEndian.Uint32(len)
			}

			// 判断缓冲区剩余数据是否足够读取一个完整的包
			if int(contentLen) > buf.Len() || contentLen == 0 {
				break
			}

			data := make([]byte, contentLen)
			_, err = buf.Read(data)

			ss := string(data)
//			if UUID > 0 {
				funRB(conn, ss, UUID)
//			} else {
//				ssList := strings.Split(ss, ",");
//				if ssList[0] == "_RegistUUID" {
//					value, _ := strconv.Atoi(ssList[1]);
//					UUID = uint64(value)
//					fmt.Println("_RegistUUID : ", UUID)
//				} else {
//					funRB(conn, ss, UUID)
//				}
//			}
			contentLen = 0
		}
	}
}

