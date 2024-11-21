package rate_limiter

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"time"
	"ych/vgo/pkg/response"
)

// Limiter 限流中间件
func Limiter(maxRequests int, timeWindow time.Duration) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(timeWindow/time.Duration(maxRequests)), maxRequests)
	return func(c *gin.Context) {
		if limiter.Allow() == false {
			response.TooManyRequests(c, "请求过于频繁~", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
