package main

import (
	"github.com/gin-gonic/gin"
	"im/internal/api/middleware"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middleware.CorsHandler())

}
