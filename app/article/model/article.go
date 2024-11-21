package model

import "ych/vgo/app/common/model"

type Article struct {
	model.BaseModel
	Title   string `gorm:"type:varchar(255);not null;column:title;default:'';comment:标题;" validate:"required,min=2" json:"title"`
	Status  int    `gorm:"column:status;type:smallint;not null;default:1;comment:状态 (1启用 2禁用)" json:"status"`
	Content string `gorm:"type:text;column:content;comment:公告内容;" json:"content"`
}
