package global

import (
	"myblog-server-go/util/timer"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"myblog-server-go/config"
)

var (
	BlogDb                 *gorm.DB
	BlogRedis              redis.UniversalClient
	BlogMongo              *qmgo.QmgoClient
	BlogConfig             config.Server
	BlogVp                 *viper.Viper
	BlogLog                *zap.Logger
	BlogTimer              timer.Timer = timer.NewTimerTask()
	BlogConcurrencyControl             = &singleflight.Group{}
	BlogRouters            gin.RoutesInfo
	BlogActiveDbname       *string
	BlackCache             local_cache.Cache
	lock                   sync.RWMutex
)
