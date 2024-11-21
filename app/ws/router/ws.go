package router

import (
	"ych/vgo/app/ws/api"
	"ych/vgo/internal/global"
)

func CollectRoutes() {
	global.Engine.GET("/ws/link", api.Link)
	global.Engine.POST("/ws/send", api.Send)
	global.Engine.POST("/ws/send_to_all", api.SendToAll)
}
