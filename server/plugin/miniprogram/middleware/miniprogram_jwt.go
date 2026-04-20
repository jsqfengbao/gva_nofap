package middleware

import (
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/miniprogram/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// MiniprogramJWTAuth 小程序JWT认证中间件
func MiniprogramJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization header
		token := c.GetHeader("Authorization")
		if token == "" {
			// 也尝试从x-token头获取
			token = c.GetHeader("x-token")
		}

		if token == "" {
			response.FailWithMessage("未提供认证token", c)
			c.Abort()
			return
		}

		// 移除Bearer前缀
		if strings.HasPrefix(token, "Bearer ") {
			token = token[7:]
		}

		// 记录认证尝试
		clientIP := c.ClientIP()
		userAgent := c.GetHeader("User-Agent")

		// 验证token
		authService := service.ServiceGroupApp.AuthService
		userID, err := authService.GetUserIDFromToken(token)
		if err != nil {
			// 记录认证失败
			global.GVA_LOG.Warn("JWT认证失败",
				zap.String("ip", clientIP),
				zap.String("user_agent", userAgent),
				zap.String("path", c.Request.URL.Path),
				zap.Error(err))

			// 检查是否需要限制IP
			if global.GVA_REDIS != nil {
				key := "auth_fail:" + clientIP
				count, _ := global.GVA_REDIS.Incr(c, key).Result()
				global.GVA_REDIS.Expire(c, key, 1*time.Hour)

				if count > 10 { // 1小时内失败超过10次
					response.FailWithMessage("认证失败次数过多，请稍后重试", c)
					c.Abort()
					return
				}
			}

			response.FailWithMessage("token验证失败", c)
			c.Abort()
			return
		}

		// 检查token是否在黑名单中
		if global.GVA_REDIS != nil {
			blacklistKey := "token_blacklist:" + token
			exists := global.GVA_REDIS.Exists(c, blacklistKey).Val()
			if exists > 0 {
				response.FailWithMessage("token已失效", c)
				c.Abort()
				return
			}
		}

		// 将用户ID和token存储到上下文
		c.Set("userID", userID)
		c.Set("token", token)
		c.Next()
	}
}

// GetUserID 从上下文获取用户ID
func GetUserID(c *gin.Context) uint {
	if userID, exists := c.Get("userID"); exists {
		if uid, ok := userID.(uint); ok {
			return uid
		}
	}
	return 0
}