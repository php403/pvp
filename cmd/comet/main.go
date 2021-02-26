package main

import (
	"flag"
	"fmt"
	log "github.com/golang/glog"
	"net"
	"os"
	"os/signal"
	"pvp/internal/comet"
	"pvp/internal/comet/conf"
	"runtime"
	"syscall"
	"time"
)



func main()  {
	flag.Parse()
	conf.Init()
	srv := comet.NewServer(conf.Conf)
	srv.AddRouter(&PingRouter{})
	if err := comet.InitTCP(srv,runtime.NumCPU());err != nil{
		panic(err)
	}
	clientTest()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("comet get a signal %s", s.String())
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

type PingRouter struct {
	comet.BaseRouter //一定要先基础BaseRouter
}

func (this *PingRouter) AfterHandle(request comet.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err !=nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (this *PingRouter) PreHandle(request comet.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err !=nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (this *PingRouter) Handle(request comet.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err !=nil {
		fmt.Println("call back ping ping ping error")
	}
}

func clientTest()  {
	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn,err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err,"client start err, exit!")
		return
	}

	for {
		_, err := conn.Write([]byte("Zinx V0.3"))
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error ")
			return
		}

		fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}