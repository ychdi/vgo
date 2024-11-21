package router

import (
	"ych/vgo/app/upload/api/ali-oss"
	"ych/vgo/app/upload/api/auto"
	"ych/vgo/app/upload/api/local"
	"ych/vgo/app/upload/api/tencent-cos"
	"ych/vgo/internal/global"
)

func CollectRoutes() {
	global.BackendRouter.POST("/upload/local/img", local.ImgUpload)
	global.BackendRouter.POST("/upload/local/video", local.VideoUpload)
	global.BackendRouter.POST("/upload/oss", ali_oss.Run)
	global.BackendRouter.POST("/upload/cos", tencent_cos.Run)
	global.BackendRouter.POST("/upload/auto", auto.Run)
}
