package router

import (
	"kube-sky/app/system/apis"

	"github.com/gin-gonic/gin"
)

func LoginLogRouter(g *gin.RouterGroup) {
	router := g.Group("/login-log")
	{
		router.GET("", apis.LoginLogList)
		router.DELETE("/:id", apis.DeleteLoginLog)
	}
}
