package bootstrap

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
	"ych/vgo/app"
	"ych/vgo/internal/global"
	"ych/vgo/internal/pkg/middleware/request-logger"
	"ych/vgo/pkg/response"
)

// Run 启动
func Run() {
	appConfig := global.Config.App
	if appConfig.Env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	global.Engine = gin.Default()

	if appConfig.RequestLog == 1 {
		global.Engine.Use(request_logger.GetLogger())
	}

	// 跨域处理
	origins := global.Config.App.ApiOrigins
	if origins == "" {
		log.Println("警告：未配置ApiOrigins，默认允许所有来源")
	}
	allowedOrigins := strings.Split(origins, ",")
	corsConfig := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	global.Engine.Use(cors.New(corsConfig))

	// 找不到路由
	global.Engine.NoRoute(func(c *gin.Context) {
		response.Fail(c, "请求地址不存在！", nil)
	})

	// 找不到方法
	global.Engine.NoMethod(func(c *gin.Context) {
		response.Fail(c, "请求方法不存在！", nil)
	})

	// 静态资源
	global.Engine.Static("/storage", "storage")

	// 初始化路由
	app.InitRouter()

	err := global.Engine.Run(appConfig.Host + ":" + appConfig.Port)
	if err != nil {
		log.Fatal(err)
		return
	}
}
