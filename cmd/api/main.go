package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"im/internal/api/middleware"
	"im/internal/api/router"
	"im/pkg/common/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Port = "7777"

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middleware.CorsHandler())
	// 注册路由
	router.Register(r)
	// 启动gin服务,不使用Run方法，以方便优雅推出
	server := &http.Server{Addr: fmt.Sprintf(":" + Port), Handler: r}
	log.Infof("Gin Listening and serving HTTP on %s", Port)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("Gin Server Error: %s", err)
		}
	}()
	// 优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Error("Gin Server Exit With TimeOut")
	}
	log.Info("Gin Server Exit")

}
