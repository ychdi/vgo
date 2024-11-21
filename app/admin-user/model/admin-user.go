package model

import "ych/vgo/app/common/model"

// AdminUser 管理员用户
type AdminUser struct {
	model.BaseModel
	UserName       string `gorm:"column:username;unique;index;default:'';type:varchar(20);not null;comment:用户名" validate:"required" form:"username" json:"username"`
	Password       string `gorm:"column:password;default:'';not null;comment:密码" form:"password" validate:"required" json:"password"`
	UserType       string `gorm:"column:user_type;default:'100';not null;comment:用户类型：(100系统用户)" form:"user_type" json:"user_type"`
	Nickname       string `gorm:"column:nickname;default:'';type:varchar(30);not null;comment:用户昵称" form:"nickname" json:"nickname"`
	Phone          string `gorm:"column:phone;default:'';type:varchar(11);not null;comment:手机" form:"phone" json:"phone"`
	Email          string `gorm:"column:email;default:'';type:varchar(50);not null;comment:用户邮箱" form:"email" json:"email"`
	Avatar         string `gorm:"column:avatar;default:'';type:varchar(255);not null;comment:用户头像" form:"avatar" json:"avatar"`
	Signed         string `gorm:"column:signed;default:'';type:varchar(255);not null;comment:个人签名" form:"signed" json:"signed"`
	Dashboard      string `gorm:"column:dashboard;default:'';type:varchar(100);not null;comment:后台首页类型" form:"dashboard" json:"dashboard"`
	Status         int    `gorm:"column:status;default:1;not null;comment:状态 (1正常 2停用)" form:"status" json:"status"`
	Super          int    `gorm:"column:super;default:2;not null;comment:是否超级管理员 (1是 2否)" form:"super" json:"super"`
	LoginIP        string `gorm:"column:login_ip;default:'';type:varchar(45);not null;comment:最后登陆IP" form:"login_ip" json:"login_ip"`
	BackendSetting string `gorm:"column:backend_setting;default:'';type:varchar(500);not null;comment:后台设置数据" form:"backend_setting" json:"backend_setting"`
	Remark         string `gorm:"column:remark;default:'';type:varchar(255);not null;comment:备注" form:"remark" json:"remark"`
}
