package Test

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"vgo/core/db"
)

func Index(ctx *gin.Context) {
	//res := time.Duration(global.App.Config.JwtConf.Timeout)
	//response.Success(ctx, "成功", map[string]interface{}{}, res)

	//query := db.GetCon().Model(&Product.Product{}).Find(&Product.Product{})
	//response.Success(ctx, "查询成功", query, nil)

	//err := db.Con().AutoMigrate(&Model.Notice{})
	//if err != nil {
	//	return
	//}
	InitCasbin()
}

func InitCasbin() {
	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && r.act == p.act
		`
	m, err := model.NewModelFromString(text)
	if err != nil {
		return
	}
	adapter, _ := gormadapter.NewAdapterByDB(db.Con())
	_, _ = casbin.NewSyncedCachedEnforcer(m, adapter)
}
