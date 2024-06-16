package controller

import (
	"time"
	dc "user_serv/db"
	"user_serv/proto"
	"user_serv/utils"

	"github.com/patrickmn/go-cache"

	"context"
)

var c = cache.New(time.Second*30, time.Second*60)

type CustomerHandler struct{}

func (*CustomerHandler) CustomerVerify(ctx context.Context, in *proto.CustomerEmailRequest, out *proto.CustomerResponse) error {
	customer := dc.Customer{}
	dc.Db.Where("Email = ?", in.Email).First(customer)
	if customer.Email != "" {
		out.Code = 500
		out.Msg = "Email has been register"
		return nil
	}

	if _, found := c.Get(in.Email); found {
		out.Code = 429
		out.Msg = "Too many requests"
		return nil
	}

	code := utils.GenRandNum(4)
	c.Set(in.Email, code, cache.DefaultExpiration)
	err := utils.SendVerifyMail([]byte(in.Email), []byte(code))
	if err != nil {
		out.Code = 502
		out.Msg = err.Error()
	} else {
		out.Code = 200
		out.Msg = "Please check email"
	}

	return nil
}

func (*CustomerHandler) CustomerRegister(ctx context.Context, in *proto.CustomerRegisterRequest, out *proto.CustomerResponse) error {
	cacheCode, ok := c.Get(in.Email)
	if !ok || cacheCode.(string) != in.Code {
		out.Code = 500
		out.Msg = "Verify failed"
		return nil
	}

	if in.Password != in.Repassword {
		out.Code = 500
		out.Msg = "Check password"
		return nil
	}

	md5Pwd := utils.Md5Sum(in.Password)

	customer := dc.Customer{
		Email:    in.Email,
		Password: md5Pwd,
		Status:   1,
	}
	result := dc.Db.Create(&customer)
	if result.Error != nil {
		out.Code = 500
		out.Msg = result.Error.Error()
		return nil
	}
	out.Code = 200
	out.Msg = "Register success"
	return nil
}

func (*CustomerHandler) CustomerLogin(ctx context.Context, in *proto.CustomerLoginRequest, out *proto.CustomerResponse) error {
	md5Pwd := utils.Md5Sum(in.Password)
	customer := dc.Customer{}

	dc.Db.Where("Email = ?", in.Email).First(&customer)
	if customer.Email != "" {
		if customer.Password == md5Pwd {
			out.Code = 200
			out.Msg = "Login in"
			out.Email = customer.Email
			return nil
		}
	}
	out.Code = 500
	out.Msg = "Login failed"
	return nil
}

func (*CustomerHandler) CustomerLogout(ctx context.Context, in *proto.CustomerLogoutRequest, out *proto.CustomerResponse) error {
	return nil
}
