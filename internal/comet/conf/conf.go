package conf

import (
	"flag"
	"github.com/spf13/viper"
	"path"
	"strings"
)

var (
	confPath  string
	//config *viper.Viper
	Conf *Config
)

type Config struct {
	Tcp *Tcp
}

type Tcp struct {
	Bind         string
	Sndbuf       int
	Rcvbuf       int
	KeepAlive    bool
	Reader       int
	ReadBuf      int
	ReadBufSize  int
	Writer       int
	WriteBuf     int
	WriteBufSize int
}

func init()  {
	flag.StringVar(&confPath, "conf", "comet.yaml", "default config path.")
}

func Init() {
	Conf = Default()
	return
}

func Default() *Config {
	viper.SetConfigName(strings.TrimSuffix(path.Base(confPath), path.Ext(confPath)))
	viper.SetConfigType(strings.TrimPrefix(path.Ext(confPath), "."))
	viper.AddConfigPath(".")
	if err:= viper.ReadInConfig(); err!= nil {
		panic(err)
	}
	var t Tcp
	if err := viper.UnmarshalKey("tcp",&t);err != nil{
		panic(err)
	}
	return &Config{
		Tcp: &t,
	}
}