package server

import (
	"fmt"
	conf "gcwguide/config"
	"github.com/gin-gonic/gin"
	"os"
)

var App *conf.Config

func init() {
	s, err := conf.InitConfig()
	if err != nil {
		fmt.Println("init server config failed, error: ", err)
		os.Exit(1)
	}
	App = s
}

func Run() {
	s := gin.Default()
	Router(s)
	s.Run(App.Server.Port)
}
