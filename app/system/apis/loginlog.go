package apis

import (
	"kube-sky/app/system/models"
	"kube-sky/pkg/conn"
	"kube-sky/pkg/pagination"
	"kube-sky/pkg/tools/response"

	"github.com/gin-gonic/gin"
)

// LoginLogList 登陆日志
func LoginLogList(c *gin.Context) {
	var (
		err    error
		list   []*models.LoginLog
		result interface{}
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: conn.Orm.Model(&models.LoginLog{}),
	}, &list, SearchParams)
	if err != nil {
		response.Error(c, err, response.LoginLogListError)
		return
	}
	response.OK(c, result, "")
}

// DeleteLoginLog 删除登陆
func DeleteLoginLog(c *gin.Context) {
	var (
		err        error
		loginLogId string
	)

	loginLogId = c.Param("id")

	err = conn.Orm.Delete(&models.LoginLog{}, loginLogId).Error
	if err != nil {
		response.Error(c, err, response.DeleteLoginLogError)
		return
	}

	response.OK(c, "", "")
}
