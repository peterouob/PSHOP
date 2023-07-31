package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	wd, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("could not get the os wd %s", err)
		fmt.Println(err)
	}
	Config.AddConfigPath(wd + "/config")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	if err := Config.ReadInConfig(); err != nil {
		panic(err)
	}
}
