package loginlog

import (
	"kube-sky/app/system/models"
	"kube-sky/pkg/conn"
	"kube-sky/pkg/logger"
	"kube-sky/pkg/tools"

	"github.com/mssola/user_agent"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context, username, status string) {
	ua := user_agent.New(c.Request.UserAgent())
	browserName, browserVersion := ua.Browser()

	loginLog := models.LoginLog{
		Username: username,
		Status:   status,
		IP:       tools.GetClientIP(c),
		Browser:  browserName + " " + browserVersion,
		System:   ua.OS(),
		Remark:   c.Request.UserAgent(),
	}

	err := conn.Orm.Create(&loginLog).Error
	if err != nil {
		logger.Errorf("登陆日志保存失败，错误：%v", err.Error())
	}
}
