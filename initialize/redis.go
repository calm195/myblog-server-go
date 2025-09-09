package initialize

import (
	"context"
	"myblog-server-go/config/cache"
	"myblog-server-go/global"

	"github.com/redis/go-redis/v9"
)

func initRedisClient(redisCfg cache.Redis) (redis.UniversalClient, error) {
	var client redis.UniversalClient

	client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Username: redisCfg.Username,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.BlogLog.Errorf("redis connect %s ping failed, err: %s", redisCfg.Address, err.Error())
		return nil, err
	}

	global.BlogLog.Infof("redis connect %s ping response: %s", redisCfg.Address, pong)
	return client, nil
}

func Redis() {
	// 初始化Redis
	global.BlogLog.Infof("redis init: %s", global.BlogConfig.Redis.Address)
	redisClient, err := initRedisClient(global.BlogConfig.Redis)
	if err != nil {
		panic(err)
	}
	global.BlogRedis = redisClient
}
