package router

import (
	"github.com/gin-gonic/gin"
	"im/internal/api/router/user"
	"net/http"
)

func Register(r *gin.Engine) {
	// index
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "just do it!"})
	})
	// 用户路由
	user.Register(r)
}
