package controller

import (
	"context"
	"net/http"

	"web_serv/utils"

	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/gin-gonic/gin"

	"time"
	"user_serv/proto"
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
	email := ctx.PostForm("email")
	code := ctx.PostForm("code")
	password := ctx.PostForm("pwd")
	repassword := ctx.PostForm("rpwd")

	client := grpc.NewClient()
	customerService := proto.NewCustomerService("go.micro.service.user_serv", client)
	response, err := customerService.CustomerRegister(context.TODO(), &proto.CustomerRegisterRequest{
		Email:      email,
		Code:       code,
		Password:   password,
		Repassword: repassword,
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
	})
}

func LoginCustomer(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("pwd")

	client := grpc.NewClient()
	customerService := proto.NewCustomerService("go.micro.service.user_serv", client)
	response, err := customerService.CustomerLogin(context.TODO(), &proto.CustomerLoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	if response.Code == http.StatusOK {
		jwt, err := utils.GenToken(response.Email, time.Hour, utils.CustomerSecretKey)
		if err != nil {
			ctx.JSON(500, gin.H{
				"code": 500,
				"msg":  err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  jwt,
		})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
	})
}

func TestGenToken(ctx *gin.Context) {
	email := ctx.PostForm("email")

	jwt, err := utils.GenToken(email, time.Second*30, utils.CustomerSecretKey)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  jwt,
	})
}

func TestValidToken(ctx *gin.Context) {
	jwt := ctx.PostForm("jwt")

	info, err := utils.VerifyToken(jwt, utils.CustomerSecretKey)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  info.UserName,
	})
}
