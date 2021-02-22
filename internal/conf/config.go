package conf

import (
	"github.com/spf13/viper"
)

func init()  {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../configs/")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}