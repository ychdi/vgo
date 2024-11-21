package backend

import (
	"github.com/gin-gonic/gin"
	"ych/vgo/app/common/backend"
	MenuModel "ych/vgo/app/menu/model"
	"ych/vgo/app/role/model"
	"ych/vgo/internal/global"
	"ych/vgo/pkg/helper"
	"ych/vgo/pkg/response"
)

// GetAll 获取全部
func GetAll(ctx *gin.Context) {
	var res []model.Role
	if err := global.DbCon.Order("id desc").Find(&res).Error; err != nil {
		response.Fail(ctx, "数据库查询失败", err.Error())
		return
	}
	response.Success(ctx, "成功", res, nil)
}

// SetMenu 设置菜单
func SetMenu(ctx *gin.Context) {
	var codes struct {
		ID    uint64   `json:"id"`
		Menus []uint64 `json:"menus"`
	}
	if err := helper.BindJSON(ctx, &codes); err != nil {
		response.Fail(ctx, "参数错误", err.Error(), nil)
		return
	}
	roleID := codes.ID
	// 清理原有菜单
	if err := global.DbCon.Delete(&model.RoleMenu{}, "role_id = ?", roleID).Error; err != nil {
		response.Fail(ctx, "菜单清理失败", err.Error())
		return
	}
	// 角色信息
	var role model.Role
	global.DbCon.First(&role, roleID)
	// 清理原有策略
	enforcer := global.Rbac
	_, err := enforcer.RemoveFilteredPolicy(0, role.Code)
	if err != nil {
		response.Fail(ctx, "策略清理失败", err.Error())
		return
	}
	// 写入菜单
	var dbMenus []MenuModel.Menu
	global.DbCon.Where("id IN ?", codes.Menus).Find(&dbMenus)
	menuMap := make(map[int]MenuModel.Menu)
	for _, dbMenu := range dbMenus {
		menuMap[int(dbMenu.ID)] = dbMenu
	}
	// 写入角色菜单关联和策略
	for _, menu := range codes.Menus {
		var item model.RoleMenu
		item.RoleId = roleID
		item.MenuId = menu
		dbMenu, exists := menuMap[int(menu)]
		if exists {
			// 写入策略
			if dbMenu.Api != "" {
				_, err2 := enforcer.AddPolicy(role.Code, dbMenu.Api, dbMenu.Act)
				if err2 != nil {
					response.Fail(ctx, "策略写入失败", err2.Error())
					return
				}
			}
		}
		if err := global.DbCon.Create(&item).Error; err != nil {
			response.Fail(ctx, err.Error(), err.Error())
			return
		}
	}
	response.Success(ctx, "设置成功", nil, nil)
}

// GetMenu 获取菜单
func GetMenu(ctx *gin.Context) {
	var codes struct {
		ID uint64 `json:"id"`
	}
	if err := helper.BindJSON(ctx, &codes); err != nil {
		response.Fail(ctx, "参数错误", err.Error(), nil)
		return
	}
	var res []model.RoleMenu
	if err := global.DbCon.Where("role_id = ?", codes.ID).Find(&res).Error; err != nil {
		response.Fail(ctx, "数据库查询失败", err.Error())
		return
	}
	response.Success(ctx, "成功", res, nil)
}

func RegisterRoleRoutes() {
	handler := backend.NewCRUDHandler(&model.Role{}, nil)
	global.BackendRouter.GET("/roles", handler.Index)
	global.BackendRouter.POST("/roles", handler.Create)
	global.BackendRouter.PUT("/roles", handler.Update)
	global.BackendRouter.DELETE("/roles", handler.Delete)

	global.BackendRouter.GET("/role/allDataSource", GetAll)
	global.BackendRouter.POST("/role/set/menu", SetMenu)
	global.BackendRouter.POST("/role/get/menu", GetMenu)
}
