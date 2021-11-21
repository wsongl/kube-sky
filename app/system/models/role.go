package models

import "kube-sky/common/models"

type Role struct {
	Key    string `gorm:"column:key;type:varchar(100);comment:标识" json:"key" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(100);comment:名称" json:"name" binding:"required"`
	Status bool   `gorm:"column:status;type:boolean;comment:状态" json:"status"`
	Remark string `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	models.BaseModel
}

func (Role) TableName() string {
	return "system_role"
}
