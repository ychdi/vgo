package permission

import (
	"fmt"
	"github.com/casbin/casbin/v2/model"
	rds "github.com/redis/go-redis/v9"
	"log"
	"ych/vgo/internal/global"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	rediswatcher "github.com/casbin/redis-watcher/v2"
)

// SetupCasbin 初始化Casbin
func SetupCasbin() *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(global.DbCon)
	if err != nil {
		panic(err)
	}
	m, err := model.NewModelFromString(`
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
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil
	}
	// 初始化Redis Watcher
	redisConf := global.Config.RedisConf
	addr := fmt.Sprintf("%v:%v", redisConf.Hostname, redisConf.HostPort)
	watcher, err := rediswatcher.NewWatcher(addr, rediswatcher.WatcherOptions{
		Channel: "/vgo-casbin", // 频道名称
		Options: rds.Options{
			Username: redisConf.UserName, // 用户名
			Password: redisConf.Password, // 密码
		},
	})

	// 设置Watcher
	err = enforcer.SetWatcher(watcher)

	// 设置回调函数，当策略发生变化时重新加载策略
	err = watcher.SetUpdateCallback(func(string) {
		err = enforcer.LoadPolicy()
	})

	// 加载策略
	err = enforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}
	return enforcer
}
