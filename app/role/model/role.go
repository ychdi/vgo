package model

import (
	"ych/vgo/app/common/model"
)

// Role 角色
type Role struct {
	model.BaseModel
	Name string `gorm:"column:name;type:varchar(50);not null;default:'';comment:角色名称" json:"name"`
	Code string `gorm:"column:code;type:varchar(50);unique;not null;default:'';comment:角色代码" json:"code"`
}
