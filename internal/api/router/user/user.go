package user

import (
	"github.com/gin-gonic/gin"
	"im/internal/api/handle/user"
)

func Register(r *gin.Engine) {
	userGroup := r.Group("/user")
	// 用户登录
	userGroup.POST("login", user.Login)
	// 用户注册
	userGroup.POST("register", user.Register)
}
