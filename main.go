package main

import (
	"fmt"
	"irita-api/common"
	"irita-api/conf"
	_ "irita-api/ctrl"
	"irita-api/gin"
	"irita-api/irita"
	"irita-api/seelog"
)

func main() {
	defer seelog.Flush()
	seelog.Infof("start irita-api")

	if err := conf.InitConfig(); err != nil {
		seelog.Errorf("init config error, %v", err)
		return
	}

	if err := irita.InitClient(); err != nil {
		seelog.Errorf("init irita sdk client error, %v", err)
		return
	}

	router := gin.CreateGin()
	common.RouterRegister.SetRouter(router)
	common.RouterRegister.InitRouter()
	if err := router.Run(fmt.Sprintf(":%d", conf.Conf.ServerConf.Port)); err != nil {
		seelog.Errorf("router run error, %v", err)
		return
	}

	seelog.Infof("stop irita-api")
}
