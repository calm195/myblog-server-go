package core

import (
	"myblog-server-go/global"
	"myblog-server-go/initialize"
	"strconv"
	"time"
)

func RunServer() {
	if global.BlogConfig.System.UseRedis {
		global.BlogLog.Info("use redis...")
		// 初始化redis服务
		initialize.Redis()
	}

	if global.BlogConfig.System.UseMongo {
		global.BlogLog.Info("use mongo...")
		// 初始化Mongo服务
		err := initialize.Mongo.Initialization()
		if err != nil {
			global.BlogLog.Panicf("mongo initialization failed: %s", err.Error())
		}
	}

	Router := initialize.Routers()

	address := ":" + strconv.Itoa(global.BlogConfig.System.Port)

	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
