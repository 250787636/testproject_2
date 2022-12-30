package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

type JiaguPolicy struct {
	Id           int    `json:"id"`             //策略id
	Status       string `json:"status"`         // 策略状态
	Name         string `json:"name"`           // 策略名称
	Template     string `json:"Template"`       // 策略配置
	Type         string `json:"type"`           // 加固策略类型
	LicKeyHelper string `json:"lic_key_helper"` // 加固策略类型
}
type Application struct {
	gorm.Model
	AppName     string `json:"app_name"`      // 系统编号
	ModelName   string `json:"model_name"`    // 模板编号
	AppCnName   string `json:"app_cn_name"`   // 系统中文全称
	AppVersion  string `json:"app_version"`   // 系统英文简称
	ModelCnName string `json:"model_cn_name"` // 模板中文全称

	RecommendPolicy int    `json:"recommend_policy"` // 推荐策略
	AppTypeID       int    `json:"app_type_id"`      // 系统类型编号
	UseUser         string `json:"use_user"`         // 使用的用户
	LastChangeTime  string `json:"last_change_time"` // 录入时间

	TheApp   string `json:"the_app"`   // 系统编号 + 系统中文全称
	TheModel string `json:"the_model"` //  模板编号 + 模板中文全称
}

func Test(t *testing.T) {
	db, err := gorm.Open("mysql", "root:www@admin@2020@tcp(172.16.102.66:33060)/middlegroundabc?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	var appTypeList Application
	db.SingularTable(true)

	if err := db.Model(Application{}).Where("id = ?", 4).First(&appTypeList).Error; err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
