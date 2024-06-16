package router

import (
	"web_serv/router"

	"github.com/gin-gonic/gin"
)

type user struct{}

func (user) Router(router *gin.RouterGroup) {
	// 发送邮件验证码
	router.POST("/send_mail", SendEmail)
	// 注册
	router.POST("/register_customer", Register)

	// 测试生成token
	router.GET("/testGenToken", TestGenToken)
	// 测试验证token
	router.GET("/testValidToken", TestValidToken)
}
