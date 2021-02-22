package comet

import "pvp/internal/comet/conf"

type Server struct {
	c *conf.Config
	//服务器的名称
	name string
	//服务绑定的端口
	Port int
}

type IServer interface{
	Start()
	Stop()
	Serve()
}

func NewServer(c *conf.Config) *Server {
	return &Server{
		c: c,
		name: "pvp",
	}
}
