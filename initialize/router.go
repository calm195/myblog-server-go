package initialize

import (
	"myblog-server-go/global"
	"myblog-server-go/router"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// 初始化总路由

func Routers() *gin.Engine {
	global.BlogLog.Info("router register start...")
	Router := gin.New()
	global.BlogLog.Info("current gin mode: " + gin.Mode())
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	systemRouter := router.GroupApp.SystemRouterGroup

	// Router.StaticFS(global.BlogConfig.Local.StorePath, justFilesFilesystem{http.Dir(global.GVA_CONFIG.Local.StorePath)}) // Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// global.GVA_LOG.Info("use middleware cors")

	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.BlogConfig.System.RouterPrefix)
	PrivateGroup := Router.Group(global.BlogConfig.System.RouterPrefix)

	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitVisitRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}

	// 注册业务路由
	initBizRouter(PrivateGroup, PublicGroup)

	global.BlogRouters = Router.Routes()

	for _, r := range global.BlogRouters {
		global.BlogLog.Infof("route: method=%s, path=%s, handler=%s", r.Method, r.Path, r.Handler)
	}

	global.BlogLog.Info("router register success")
	return Router
}

// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.GroupApp
}

func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]

	holder(publicGroup, privateGroup)
}
