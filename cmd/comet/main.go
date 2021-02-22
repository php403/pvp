package main

import (
	"flag"
	log "github.com/golang/glog"
	"os"
	"os/signal"
	"pvp/internal/comet"
	"pvp/internal/comet/conf"
	"runtime"
	"syscall"
)

func main()  {
	flag.Parse()
	conf.Init()
	srv := comet.NewServer(conf.Conf)
	if err := comet.InitTCP(srv,runtime.NumCPU());err != nil{
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("goim-comet get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//srv.Close()
			log.Infof("goim-comet [version] exit")
			log.Flush()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

/*func broadcaster() {

}

func handleConn(conn net.Conn) {
	defer conn.Close()

}*/
