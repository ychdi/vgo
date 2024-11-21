package auth

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"ych/vgo/internal/global"
	"ych/vgo/pkg/response"
)

// AdminAuthMiddleware JWT验证中间件
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			// Header format is invalid
			response.NotLogin(c, "Token无效", nil)
			c.Abort()
			return
		}
		tokenString := parts[1]
		if tokenString == "" {
			// Header is missing
			response.NotLogin(c, "Token无效", nil)
			c.Abort()
			return
		}
		claims, err := ParseAdminToken(tokenString)
		if err != nil {
			//解析Fail
			response.NotLogin(c, "Token无效", nil)
			c.Abort()
			return
		}

		if AssertIsTokenInvalid(c, tokenString) {
			//黑名单
			response.NotLogin(c, "Token无效", nil)
			c.Abort()
			return
		}
		JwtConf := global.Config.JwtConf
		redisCon := global.RedisCon
		// 是否单设备登录
		if JwtConf.AdminSingleLogin == 1 {
			redisToken := redisCon.Get(c, "admin_token"+strconv.Itoa(int(claims.UserID)))
			if redisToken == nil || redisToken.Val() != tokenString {
				// Token is invalid --backend-- redis00
				response.NotLogin(c, "Token无效", nil)
				c.Abort()
				return
			}
		} else {
			redisToken := redisCon.LRange(c, "admin_token_list"+strconv.Itoa(int(claims.UserID)), 0, -1)
			if redisToken == nil || len(redisToken.Val()) == 0 {
				response.NotLogin(c, "Token无效", nil)
				c.Abort()
				return
			}
			if !strings.Contains(strings.Join(redisToken.Val(), ""), tokenString) {
				response.NotLogin(c, "Token无效", nil)
				c.Abort()
				return
			}
		}
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Set("super", claims.Super)
		c.Next()
	}
}

// UserAuthMiddleware 前台用户的JWT验证中间件
func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			// Header format is invalid
			response.NotLogin(c, "Token无效", nil)
			c.Abort()
			return
		}
		tokenString := parts[1]
		if tokenString == "" {
			// Header is missing
			response.NotLogin(c, "Token无效", nil)
			c.Abort()
			return
		}

		claims, err := ParseUserToken(tokenString)
		if err != nil {
			// 解析Fail
			response.NotLogin(c, "Token无效", nil)
			c.Abort()
			return
		}

		if AssertApiTokenIsInvalid(c, tokenString) {
			// 黑名单
			response.NotLogin(c, "Token无效", nil)
			c.Abort()
			return
		}

		JwtConf := global.Config.JwtConf
		redisCon := global.RedisCon
		// 是否单设备登录
		if JwtConf.ApiSingleLogin == 1 {
			redisToken := redisCon.Get(c, "api_token"+strconv.FormatInt(int64(claims.UserID), 10))
			if redisToken == nil || redisToken.Val() != tokenString {
				//--api-- redis00
				response.NotLogin(c, "Token无效", nil)
				c.Abort()
				return
			}
		} else {
			redisToken := redisCon.LRange(c, "api_token_list"+strconv.FormatInt(int64(claims.UserID), 10), 0, -1)
			if redisToken == nil || len(redisToken.Val()) == 0 {
				//--api-- redis0
				response.NotLogin(c, "Token无效", nil)
				c.Abort()
				return
			}
			if !strings.Contains(strings.Join(redisToken.Val(), ""), tokenString) {
				// --api-- redis02
				response.NotLogin(c, "Token无效", nil)
				c.Abort()
				return
			}
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
