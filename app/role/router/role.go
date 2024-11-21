package router

import (
	"ych/vgo/app/role/backend"
)

func CollectRoutes() {
	backend.RegisterRoleRoutes()
}
