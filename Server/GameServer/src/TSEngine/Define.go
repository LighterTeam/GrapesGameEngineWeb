package TSEngine

import (
	"net"
)

type MessagePacket struct {
	T int32 //Opcode
	S uint64 //SendToUuid
	D string //Data
}

type SpriteBullet struct {
	Master int32
	Handle int32
	X float32
	Y float32
}

type Sprite struct {
	Handle int32
	X float32
	Y float32
}

type Client struct {
	Uuid uint64
	Spt Sprite
	Conn net.Conn
}
