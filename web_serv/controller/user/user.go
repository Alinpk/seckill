package controller

import (
	"context"
	"net/http"

	"web_serv/utils"

	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/gin-gonic/gin"

	"user_serv/proto"

	"fmt"
)

func SendEmail(ctx *gin.Context) {
	email := ctx.PostForm("email")

	if ok := utils.VerifyEmail(email); ok {
		client := grpc.NewClient()

		customerService := proto.NewCustomerService("go.micro.service.user_serv", client)
		response, err := customerService.CustomerVerify(context.TODO(), &proto.CustomerEmailRequest{
			Email: email,
		})

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  err.Error(),
			})
			fmt.Println("err.Error():", err.Error())
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": response.Code,
				"msg":  response.Msg,
			})
		}
		// fmt.Println("Code", response.Code, "\n", "msg:", response.Msg)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  "Invalid mail format",
	})
}

func RegisterCustomer(ctx *gin.Context) {

}

func LoginCustomer(ctx *gin.Context) {

}

func TestGenToken(ctx *gin.Context) {

}

func TestValidToken(ctx *gin.Context) {

}
