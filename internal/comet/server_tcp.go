package comet

import (
	"errors"
	"fmt"
	log "github.com/golang/glog"
	"net"
	_ "pvp/internal/comet/conf"
)

func init() {

}

func InitTCP(server *Server,accept int)(err error) {
	var (
		bind     string
		listener *net.TCPListener
		addr     *net.TCPAddr
	)
	bind = server.c.Tcp.Bind
	if addr, err = net.ResolveTCPAddr("tcp", bind); err != nil {
		log.Errorf("net.ResolveTCPAddr(tcp, %s) error(%v)", bind, err)
		return
	}
	if listener, err = net.ListenTCP("tcp", addr); err != nil {
		log.Errorf("net.ListenTCP(tcp, %s) error(%v)", bind, err)
		return
	}
	log.Infof("start tcp listen: %s", bind)
	// split N core accept
	for i := 0; i < accept; i++ {
		go acceptTCP(server, listener)
	}
	return
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	//回显业务
	if _, err := conn.Write(data[:cnt]); err !=nil {
		return errors.New("CallBackToClient error")
	}
	return nil
}

func acceptTCP(server *Server, lis *net.TCPListener) {
	var (
		conn *net.TCPConn
		err  error
	)
	for {
		if conn, err = lis.AcceptTCP(); err != nil {
			// if listener close then return
			log.Errorf("listener.Accept(\"%s\") error(%v)", lis.Addr().String(), err)
			return
		}
		if err = conn.SetKeepAlive(server.c.Tcp.KeepAlive); err != nil {
			log.Errorf("conn.SetKeepAlive() error(%v)", err)
			return
		}
		if err = conn.SetReadBuffer(server.c.Tcp.Rcvbuf); err != nil {
			log.Errorf("conn.SetReadBuffer() error(%v)", err)
			return
		}
		if err = conn.SetWriteBuffer(server.c.Tcp.Sndbuf); err != nil {
			log.Errorf("conn.SetWriteBuffer() error(%v)", err)
			return
		}
		var cid uint32
		cid = 0
		go func () {
			//不断的循环从客户端获取数据
			for  {
				buf := make([]byte, server.c.Tcp.Rcvbuf)
				cnt, err := conn.Read(buf)
				if err != nil {
					fmt.Println("recv buf err ", err)
					continue
				}
				fmt.Println("cnt data :",cnt)

				dealConn := NewConnection(conn,cid,CallBackToClient,server.Router)
				cid ++

				//3.4 启动当前链接的处理业务
				go dealConn.Start()
			}
		}()
	}
}

func (s *Server)AddRouter(router IRouter) {
	s.Router = router
	fmt.Println("Add Router succ! " )
}





