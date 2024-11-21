package model

import (
	"strconv"
	"ych/vgo/app/common/model"
	"ych/vgo/internal/global"
)

// RoleMenu 角色菜单
type RoleMenu struct {
	model.BaseModel
	RoleId uint64 `gorm:"column:role_id;not null;default:0;comment:角色ID" json:"role_id"`
	MenuId uint64 `gorm:"column:menu_id;not null;default:0;comment:菜单ID" json:"menu_id"`
}

// GetMenuIdsByRoleId 获取角色菜单ID列表
func GetMenuIdsByRoleId(userID uint64) []uint64 {
	// 获取角色
	enforcer := global.Rbac
	roles, err := enforcer.GetRolesForUser(strconv.FormatUint(userID, 10))
	if err != nil {
		return nil
	}
	var dbRoles []Role
	global.DbCon.Find(&dbRoles, "code in (?)", roles)
	roleIDs := make([]uint64, 0)
	for _, rid := range dbRoles {
		roleIDs = append(roleIDs, rid.ID)
	}
	// 获取角色菜单关联
	var roleMenu []RoleMenu
	global.DbCon.Find(&roleMenu, "role_id in (?)", roleIDs)
	menuIDs := make([]uint64, 0)
	for _, roleMenu := range roleMenu {
		menuIDs = append(menuIDs, roleMenu.MenuId)
	}
	return menuIDs
}
