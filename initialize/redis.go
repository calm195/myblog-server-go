package initialize

import (
	"context"
	"myblog-server-go/config/cache"
	"myblog-server-go/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func initRedisClient(redisCfg cache.Redis) (redis.UniversalClient, error) {
	var client redis.UniversalClient
	// 使用单例模式
	client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.BlogLog.Error("redis connect ping failed, err:", zap.String("Address", redisCfg.Address), zap.Error(err))
		return nil, err
	}

	global.BlogLog.Info("redis connect ping response:", zap.String("Address", redisCfg.Address), zap.String("pong", pong))
	return client, nil
}

func Redis() {
	redisClient, err := initRedisClient(global.BlogConfig.Redis)
	if err != nil {
		panic(err)
	}
	global.BlogRedis = redisClient
}
