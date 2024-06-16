package controller

import (
	"github.com/gin-gonic/gin"
)

type user struct{}

var User user

func (user) Router(router *gin.RouterGroup) {
	// 发送邮件验证码
	router.POST("/send_mail", SendEmail)
	// 注册
	router.POST("/register_customer", RegisterCustomer)

	router.POST("/login_customer", LoginCustomer)

	// 测试生成token
	router.GET("/testGenToken", TestGenToken)

	// 测试验证token
	router.GET("/testValidToken", TestValidToken)
}
