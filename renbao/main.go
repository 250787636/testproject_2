package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/domain/ngiam-rst/cas/serviceValidate", serviceValidate)
	r.Run(":8888")
}

func serviceValidate(c *gin.Context) {
	service := c.Query("service")
	ticket := c.Query("ticket")
	myJson := c.Query("format")
	fmt.Println(service)
	fmt.Println(ticket)
	fmt.Println(myJson)

	if ticket == "1d23456" && service == "https://127.0.0.1:9100/ccid/handleAuthFromMainPlatform" &&
		myJson == "json" {
		c.JSON(http.StatusOK, gin.H{ // 正确的json格式
			"serviceResponse": gin.H{
				"authenticationSuccess": gin.H{
					"user": "123456 ", //应用账户名
					"attributes": gin.H{
						"userId":     "4678-8a9d",
						"userApiKey": "9678-8ea9d",
					},
				},
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{ //失败的json格式
			"serviceResponse": gin.H{
				"authenticationFailure": gin.H{
					"code":        "INVALID_TICKET",
					"description": "Ticket ST-1856339-aA5Yuvrxzpv8Tau1cYQ7 not recognized",
				},
			},
		})
	}

}
