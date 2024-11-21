package router

import (
	"ych/vgo/app/article/api"
	"ych/vgo/app/article/backend"
)

func CollectRoutes() {
	backend.RegisterArticleRoutes()
	api.RegisterArticleRoutes()
}
