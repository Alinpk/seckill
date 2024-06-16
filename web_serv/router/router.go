package router

import (
	"github.com/gin-gonic/gin"

	"web_serv/controller/user"
)

func InitRouter(router *gin.Engine) {
	user_group := router.Group("/user")

	controller.User.Router(user_group)
}
