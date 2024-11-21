package router

import (
	"ych/vgo/app/admin-user/backend"
)

func CollectRoutes() {
	backend.RegisterAdminUserRoutes()
}
