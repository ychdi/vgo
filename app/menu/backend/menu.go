package backend

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"ych/vgo/app/common/backend"
	"ych/vgo/app/menu/model"
	Role "ych/vgo/app/role/model"
	"ych/vgo/internal/global"
	"ych/vgo/pkg/enum"
	"ych/vgo/pkg/response"
)

// Index 无限极分类菜单结构
func Index(ctx *gin.Context) {
	var menus []model.Menu
	var err error
	if ctx.GetInt("super") == enum.CommonEnable { // 超级管理员
		err = global.DbCon.Order("sort desc").Find(&menus).Error
	} else {
		menuIDs := Role.GetMenuIdsByRoleId(ctx.GetUint64("userID"))
		err = global.DbCon.Order("sort desc").Find(&menus, "id in (?)", menuIDs).Error
	}
	if err != nil {
		response.Fail(ctx, "数据库查询失败", err)
		return
	}
	menuTree := model.BuildMenuTree(menus, 0)
	response.Success(ctx, "成功", menuTree, nil)
}

// Buttons 操作按钮
func Buttons(ctx *gin.Context) {
	var data = make(map[string][]string)
	var menus []model.Menu
	var err error

	if ctx.GetInt("super") == enum.CommonEnable { // 超级管理员
		if err = global.DbCon.Where("type = ?", 1).Find(&menus).Error; err != nil {
			response.Fail(ctx, "数据库查询失败", err)
			return
		}
		for _, menu := range menus {
			var buttonMenus []model.Menu
			if err = global.DbCon.Where("parent_id = ? and type = ?", menu.ID, 2).Find(&buttonMenus).Error; err != nil {
				response.Fail(ctx, "数据库查询失败", err)
				return
			}
			if len(buttonMenus) > 0 {
				bbt := make([]string, len(buttonMenus))
				for i, buttonMenu := range buttonMenus {
					bbt[i] = buttonMenu.Name
				}
				data[menu.Name] = bbt
			}
		}
		response.Success(ctx, "成功", data, nil)
		return
	}
	menuIDs := Role.GetMenuIdsByRoleId(ctx.GetUint64("userID"))
	if err = global.DbCon.Where("type = ? and id in (?)", 1, menuIDs).Find(&menus).Error; err != nil {
		response.Fail(ctx, "数据库查询失败", err)
		return
	}
	// 获取角色
	enforcer := global.Rbac
	roles, err := enforcer.GetRolesForUser(strconv.FormatUint(ctx.GetUint64("userID"), 10))
	if err != nil {
		response.Fail(ctx, "角色获取失败", err)
		return
	}
	for _, menu := range menus {
		var buttonMenus []model.Menu
		if err = global.DbCon.Where("parent_id = ? and type = ?", menu.ID, 2).Find(&buttonMenus).Error; err != nil {
			response.Fail(ctx, "数据库查询失败", err)
			return
		}
		if len(buttonMenus) == 0 {
			continue
		}
		bbt := make([]string, 0, len(buttonMenus))
		for _, buttonMenu := range buttonMenus {
			if buttonMenu.Name == "" {
				continue
			}
			// 判断按钮是否有权限
			hasPermission := false
			for _, role := range roles {
				if policy := enforcer.HasPolicy(role, buttonMenu.Api, buttonMenu.Act); policy {
					hasPermission = true
					break
				}
			}
			if hasPermission {
				bbt = append(bbt, buttonMenu.Name)
			}
		}
		if len(bbt) > 0 {
			data[menu.Name] = bbt
		}
	}
	response.Success(ctx, "成功", data, nil)
}

// MenuSelect 下拉树结构
type MenuSelect struct {
	Value    uint64       `json:"value"`
	Label    string       `json:"label"`
	Children []MenuSelect `json:"children"`
}

// convertToMenuSelect 将Menu结构体转换为MenuSelect结构体
func convertToMenuSelect(menu model.Menu) MenuSelect {
	menuSelect := MenuSelect{
		Value:    menu.ID,    // 假设Menu结构体中有ID字段
		Label:    menu.Title, // 假设Menu结构体中有Name字段
		Children: []MenuSelect{},
	}
	for _, child := range menu.Children {
		menuSelect.Children = append(menuSelect.Children, convertToMenuSelect(child))
	}
	return menuSelect
}

// GetSelectTree 获取下拉树结构
func GetSelectTree(ctx *gin.Context) {
	var menus []model.Menu
	if err := global.DbCon.Order("sort desc").Find(&menus).Error; err != nil {
		response.Fail(ctx, "数据库查询失败", err)
		return
	}
	menuTree := model.BuildMenuTree(menus, 0)
	menuSelects := make([]MenuSelect, len(menuTree))
	for i, menu := range menuTree {
		menuSelects[i] = convertToMenuSelect(menu)
	}
	response.Success(ctx, "查询成功", menuSelects, nil)
}

func RegisterMenuRoutes() {
	handler := backend.NewCRUDHandler(&model.Menu{}, nil)
	global.BackendRouter.POST("/menus", handler.Create)
	global.BackendRouter.PUT("/menus", handler.Update)
	global.BackendRouter.DELETE("/menus", handler.Delete)

	global.BackendRouter.GET("/menus", Index)
	global.BackendRouter.GET("/buttons", Buttons)
	global.BackendRouter.GET("/menu/selectTreeDataSource", GetSelectTree)
}
