package router

import (
	"ych/vgo/app/notice/api"
	"ych/vgo/app/notice/backend"
)

func CollectRoutes() {
	backend.RegisterNoticeRoutes()
	api.RegisterNoticeApiRoutes()
}
