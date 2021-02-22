package comet

import (
	"fmt"
	log "github.com/golang/glog"
	"net"
	_ "pvp/internal/comet/conf"
)

func init()  {

}

func InitTCP(server *Server,accept int)(err error)  {
	var (
		bind     string
		listener *net.TCPListener
		addr     *net.TCPAddr
	)
	bind = server.c.Tcp.Bind
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", bind)
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
		go func () {
			//不断的循环从客户端获取数据
			for  {
				buf := make([]byte, server.c.Tcp.Rcvbuf)
				cnt, err := conn.Read(buf)
				if err != nil {
					fmt.Println("recv buf err ", err)
					continue
				}
				//回显
				if _, err := conn.Write(buf[:cnt]); err !=nil {
					fmt.Println("write back buf err ", err)
					continue
				}
			}
		}()
	}
}




