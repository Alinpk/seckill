package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	user_group := router.Group("/user")

	user.Router(user_group)
}
