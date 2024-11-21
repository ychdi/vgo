package permission

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"regexp"
	"ych/vgo/pkg/response"
)

// Check 权限检查中间件
func Check(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		super := c.GetInt("super") // 超级管理员
		if super == 1 {
			c.Next()
			return
		}
		obj := c.Request.URL.Path
		// 排除数据源权限判断
		re := regexp.MustCompile(`DataSource`)
		if re.MatchString(obj) {
			c.Next()
			return
		}
		act := c.Request.Method
		sub := c.GetStringSlice("role") // 获取角色
		for _, v := range sub {
			if ok, _ := e.Enforce(v, obj, act); ok {
				c.Next()
				return
			}
		}
		response.Forbidden(c, "无权限访问", nil)
		c.Abort()
	}
}
