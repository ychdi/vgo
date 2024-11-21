package router

import (
	"ych/vgo/app/common/api"
	"ych/vgo/internal/global"
)

func CollectRoutes() {
	global.Engine.GET("/test", api.Test)
}
