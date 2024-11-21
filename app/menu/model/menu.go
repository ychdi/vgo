package model

import (
	"gorm.io/gorm"
	"ych/vgo/app/common/model"
)

// Menu 菜单
type Menu struct {
	model.BaseModel
	ParentId    uint64                 `gorm:"column:parent_id;type:bigint;not null;default:0;comment:父ID" json:"parent_id"`
	Path        string                 `gorm:"column:path;type:varchar(50);not null;default:'';comment:路由访问路径" json:"path"`
	Name        string                 `gorm:"column:name;type:varchar(50);not null;default:'';comment:路由name" json:"name"`
	Redirect    string                 `gorm:"column:redirect;type:varchar(255);not null;default:'';comment:路由重定向地址" json:"redirect"`
	Api         string                 `gorm:"column:api;type:varchar(255);not null;default:'';comment:请求接口地址" json:"api"`
	Act         string                 `gorm:"column:act;type:varchar(255);not null;default:'';comment:请求方式" json:"act"`
	Component   string                 `gorm:"column:component;type:varchar(255);not null;default:'';comment:视图文件路径" json:"component"`
	Icon        string                 `gorm:"column:icon;type:varchar(50);not null;default:'';comment:菜单和面包屑对应的图标" json:"icon"`
	Title       string                 `gorm:"column:title;type:varchar(50);not null;default:'';comment:路由标题(菜单名称)" json:"title"`
	ActiveMenu  string                 `gorm:"column:activeMenu;type:varchar(500);not null;default:'';comment:是否在菜单中隐藏,需要高亮的path" json:"activeMenu"`
	IsLink      string                 `gorm:"column:isLink;type:varchar(500);not null;default:'';comment:路由外链时填写的访问地址" json:"isLink"`
	IsHide      int                    `gorm:"column:isHide;type:smallint;not null;default:2;comment:是否在菜单中隐藏 (1是 2否)" json:"isHide"`
	IsFull      int                    `gorm:"column:isFull;type:smallint;not null;default:2;comment:菜单是否全屏 (1是 2否)" json:"isFull"`
	IsAffix     int                    `gorm:"column:isAffix;type:smallint;not null;default:1;comment:菜单是否固定在标签页中 (1是 2否)" json:"isAffix"`
	IsKeepAlive int                    `gorm:"column:isKeepAlive;type:smallint;not null;default:2;comment:当前路由是否缓存 (1是 2否)" json:"isKeepAlive"`
	Status      int                    `gorm:"column:status;type:smallint;not null;default:1;comment:状态 (1正常 2停用)" json:"status"`
	Type        int                    `gorm:"column:type;type:smallint;not null;default:1;comment:类型 (1菜单2按钮3外链4Iframe)" json:"type"`
	Sort        int                    `gorm:"column:sort;type:int;not null;default:0;comment:排序" json:"sort"`
	Meta        map[string]interface{} `gorm:"-"  json:"meta"` // 忽略数据库字段，使用Getter方法
	Children    []Menu                 `gorm:"foreignKey:ParentId" json:"children"`
}

// AfterFind 回调函数，在查询后执行
func (m *Menu) AfterFind(_ *gorm.DB) (err error) {
	m.Meta = m.getMetaAttribute()
	return
}

// GetMetaAttribute 获取菜单元数据属性
func (m *Menu) getMetaAttribute() map[string]interface{} {
	meta := map[string]interface{}{
		"icon":        m.Icon,
		"title":       m.Title,
		"activeMenu":  m.ActiveMenu,
		"isLink":      m.IsLink,
		"isHide":      m.IsHide == 1,
		"isFull":      m.IsFull == 1,
		"isAffix":     m.IsAffix == 1,
		"isKeepAlive": m.IsKeepAlive == 1,
	}
	return meta
}

// BuildMenuTree 构建菜单树
func BuildMenuTree(menus []Menu, parentID uint64) []Menu {
	var tree []Menu
	for _, menu := range menus {
		if menu.ParentId == parentID {
			children := BuildMenuTree(menus, menu.ID)
			if len(children) > 0 {
				menu.Children = children
			}
			tree = append(tree, menu)
		}
	}
	return tree
}
