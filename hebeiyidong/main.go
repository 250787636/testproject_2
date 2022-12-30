package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ky/ssp/hebeiyidong/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.GenerationLogs)
	r.POST("/lic/list", CheckLoginIn)
	r.POST("/task/myloginin", MyLoginIn)
	r.POST("/report/normal", Normal)
	r.POST("/a/normaltype", NormalType)
	r.Run(":8888")
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"error"`
	Data interface{} `json:"data"`
}

func CheckLoginIn(c *gin.Context) {
	c.JSON(200, Response{Code: 200, Data: "成功登录"})
	logrus.Println("这是CheckLoginIn方法")
}
func MyLoginIn(c *gin.Context) {
	c.JSON(500, Response{Code: 500, Msg: "登录失败"})
	logrus.Println("这是MyLoginIn方法")
}
func Normal(c *gin.Context) {
	c.JSON(200, Response{Code: 200, Data: "普通方法成功"})
	logrus.Println("这是Normal方法")
}
func NormalType(c *gin.Context) {
	c.JSON(200, Response{Code: 500, Msg: "普通方法失败"})
	logrus.Println("这是NormalType方法")
}
