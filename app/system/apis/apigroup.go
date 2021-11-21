package apis

import (
	"kube-sky/app/system/models"
	"kube-sky/pkg/conn"
	"kube-sky/pkg/pagination"
	"kube-sky/pkg/tools/response"

	"github.com/gin-gonic/gin"
)

// ApiGroupList 接口分组列表
func ApiGroupList(c *gin.Context) {
	var (
		err          error
		apiGroupList []*models.ApiGroup
		result       interface{}
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: conn.Orm.Model(&models.ApiGroup{}).Order("id"),
	}, &apiGroupList, SearchParams)
	if err != nil {
		response.Error(c, err, response.ApiGroupListError)
		return
	}
	response.OK(c, result, "")
}

// SaveApiGroup 保存接口分组
func SaveApiGroup(c *gin.Context) {
	var (
		err           error
		apiGroup      models.ApiGroup
		apiGroupCount int64
	)

	err = c.ShouldBind(&apiGroup)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	db := conn.Orm.Model(&models.ApiGroup{})

	if apiGroup.Id != 0 {
		db = db.Where("id = ?", apiGroup.Id)
	} else {
		err = conn.Orm.Model(&models.ApiGroup{}).
			Where(`"name" = ?`, apiGroup.Name).
			Count(&apiGroupCount).Error
		if err != nil {
			response.Error(c, err, response.ApiGroupExistError)
			return
		}
	}

	err = db.Save(&apiGroup).Error
	if err != nil {
		response.Error(c, err, response.SaveApiGroupError)
		return
	}

	response.OK(c, "", "")
}

// DeleteApiGroup 删除接口分组
func DeleteApiGroup(c *gin.Context) {
	var (
		err        error
		apiCount   int64
		apiGroupId string
	)

	apiGroupId = c.Param("id")

	err = conn.Orm.Model(&models.Api{}).Where(`"group" = ?`, apiGroupId).Count(&apiCount).Error
	if err != nil {
		response.Error(c, err, response.GetApiError)
		return
	}
	if apiCount > 0 {
		response.Error(c, err, response.ApiGroupUsedError)
		return
	}

	err = conn.Orm.Delete(&models.ApiGroup{}, apiGroupId).Error
	if err != nil {
		response.Error(c, err, response.DeleteApiGroupError)
		return
	}

	response.OK(c, "", "")
}
