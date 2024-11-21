package app

import (
	AdminUserRouter "ych/vgo/app/admin-user/router"
	ArticleRouter "ych/vgo/app/article/router"
	CommonRouter "ych/vgo/app/common/router"
	MenuRouter "ych/vgo/app/menu/router"
	NoticeRouter "ych/vgo/app/notice/router"
	RoleRouter "ych/vgo/app/role/router"
	UploadRouter "ych/vgo/app/upload/router"
	wsrouter "ych/vgo/app/ws/router"

	"ych/vgo/internal/global"
	"ych/vgo/internal/pkg/middleware/auth"
	"ych/vgo/internal/pkg/middleware/permission"
)

func InitRouter() {
	global.BackendRouter = global.Engine.Group("/backend")
	global.BackendRouter.Use(auth.AdminAuthMiddleware(), permission.Check(global.Rbac))

	global.ApiRouter = global.Engine.Group("/api/v1")

	ArticleRouter.CollectRoutes()
	NoticeRouter.CollectRoutes()

	AdminUserRouter.CollectRoutes()
	CommonRouter.CollectRoutes()
	wsrouter.CollectRoutes()
	UploadRouter.CollectRoutes()

	MenuRouter.CollectRoutes()
	RoleRouter.CollectRoutes()
}
