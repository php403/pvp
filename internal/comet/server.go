package comet

import (
	"net"
	"pvp/internal/comet/conf"
)

type Server struct {
	c *conf.Config
	//服务器的名称
	name string
	//服务绑定的端口
	Port int
	Router IRouter
}

type IServer interface{
	Start()
	Stop()
	Serve()
	AddRouter(router IRouter)
}

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
}


func NewServer(c *conf.Config) *Server {
	return &Server{
		c: c,
		name: "pvp",
		Router: nil,
	}
}


type HandFunc func(*net.TCPConn, []byte, int) error

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	handleAPI HandFunc
	ExitBuffChan chan bool
	Router  IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, callbackApi HandFunc,router IRouter) *Connection{
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		handleAPI: callbackApi,
		ExitBuffChan: make(chan bool, 1),
		Router: router,
	}
	return c
}

