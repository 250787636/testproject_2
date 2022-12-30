package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/api/getUserInfoByToken", SADISso)
	r.GET("/api/appToken", SADIGetAppToken)
	r.Run(":9999")
}

//获取的信息
type appGetDetail struct {
	AppKey    string `json:"appKey" binding:"required"`
	AppSecret string `json:"appSecret" binding:"required"`
}
type userGetDetail struct {
	AppToken string `json:"appToken" binding:"required"`
	Token    string `json:"token" binding:"required"`
	SysName  string `json:"sysName" binding:"required"`
}

// 返回的信息
type appReturnDetail struct {
	AppToken   string `json:"appToken"`
	ExpiryDate int64  `json:"expiryDate"`
}
type userReturnDetail struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    struct {
		Name     string `json:"name"`
		NickName string `json:"nickName"`
		Email    string `json:"email"`
	} `json:"data"`
}

// 获取AppToken
func SADIGetAppToken(r *gin.Context) {
	var appDetail appGetDetail
	appKey := "c0521f6e-1ace-4815-883c-0cbc77e5d6ab"
	appSecret := "59df2fb5-2908-49bc-9d3d-a727302c859e"
	appDetail.AppKey = r.Query("appKey")
	appDetail.AppSecret = r.Query("appSecret")
	if appDetail.AppKey != appKey || appDetail.AppSecret != appSecret {
		r.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "appKey或者appSecret有误",
		})
		return
	}
	appReturnDetail := appReturnDetail{
		AppToken:   "72d4a54a-91ff-41a1-9de4-c5ad3cd9f642",
		ExpiryDate: 1635475552917,
	}
	r.JSON(http.StatusOK, appReturnDetail)
}

// 获取用户信息
func SADISso(r *gin.Context) {
	var userDetail userGetDetail
	if err := r.ShouldBindJSON(&userDetail); err != nil {
		r.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"success": false,
			"msg":     err.Error(),
			"data":    nil,
		})
		return
	}
	userReturnDetail := new(userReturnDetail)
	userReturnDetail.Code = 200
	userReturnDetail.Success = true
	userReturnDetail.Msg = "用户获取成功"

	if userDetail.Token == "31f0692eb40840aab95591cd1faf8be9" {
		userReturnDetail.Data.Name = "wenjng@163.com"
		userReturnDetail.Data.NickName = "nihao"
		userReturnDetail.Data.Email = "weneng@163.com"
	} else if userDetail.Token == "11f0692eb40840aab95591cd1faf8be8" {
		userReturnDetail.Data.Name = "jinyun@163.com"
		userReturnDetail.Data.NickName = "jinyun"
		userReturnDetail.Data.Email = "jinyun@163.com"
	} else if userDetail.Token == "21f0692eb40840aab95591cd1faf8be7" {
		userReturnDetail.Data.Name = "tanyong@163.com"
		userReturnDetail.Data.NickName = "tanyong"
		userReturnDetail.Data.Email = "tanyong@163.com"
	} else if userDetail.Token == "41f0692eb40840aab95591cd1faf8be6" {
		userReturnDetail.Data.Name = "yongpin@163.com"
		userReturnDetail.Data.NickName = "yongpin"
		userReturnDetail.Data.Email = "yongpin@163.com"
	} else {
		userReturnDetail.Success = false
		userReturnDetail.Msg = "用户token不存在"
	}
	r.JSON(http.StatusOK, userReturnDetail)
}
