package core

import (
	"fmt"
	"go.uber.org/zap"
	"myblog-server-go/global"
	"myblog-server-go/initialize"
	"time"
)

func RunServer() {
	if global.BlogConfig.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	if global.BlogConfig.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.BlogConfig.System.Port)

	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
