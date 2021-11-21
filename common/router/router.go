package router

import (
	"net/http"
	"kube-sky/pkg/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// 加载路由
func Setup(g *gin.Engine) {
	// 使用zap接收gin框架默认的日志并配置
	g.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 页面不存在")
	})

	// pprof router
	pprof.Register(g)

	// cors， 跨域
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowOrigins:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	g.Use(cors.New(config))

	// 路由版本
	routerVersion := g.Group("/api/v1")

	// 注册公共接口路由，非业务相关路由
	registerPublicRouter(routerVersion)

	// 注册系统管理路由
	registerSystemRouter(routerVersion)
}
