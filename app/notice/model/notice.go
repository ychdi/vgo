package model

import (
	"gorm.io/plugin/soft_delete"
	"ych/vgo/app/common/model"
)

type Notice struct {
	model.BaseModel
	Title     string                `gorm:"type:varchar(255);not null;column:title;uniqueIndex:udx_title;default:'';comment:标题;" validate:"required" json:"title"`
	Status    int                   `gorm:"column:status;type:smallint;not null;default:1;comment:状态 (1启用 2禁用)" json:"status"`
	Content   string                `gorm:"type:text;column:content;comment:公告内容;" json:"content"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_title"`
}
