package controller

import (
	dc "user_serv/db"
	"user_serv/proto"
	"user_serv/utils"

	"github.com/patrickmn/go-cache"

	"context"
)

type SellerHandler struct{}

func (SellerHandler) SellerVerify(ctx context.Context, in *proto.SellerEmailRequest, out *proto.SellerResp) error {
	if _, ok := c.Get(in.Email); ok {
		out.Code = 500
		out.Msg = "Verify too frequence"
		return nil
	}

	seller := dc.Seller{}
	dc.Db.Where("Email = ?", in.Email).First(&seller)
	if seller.Email == "" {
		out.Code = 500
		out.Msg = "Invalid seller"
		return nil
	}

	code := utils.GenRandNum(4)
	if err := utils.SendVerifyMail([]byte(in.Email), []byte(code)); err != nil {
		out.Code = 500
		out.Msg = "Email send failed:" + err.Error()
		return nil
	}

	c.Set(in.Email, code, cache.DefaultExpiration)
	out.Code = 200
	out.Msg = "Please check email"
	return nil
}

func (SellerHandler) SellerLogin(ctx context.Context, in *proto.SellerLoginRequest, out *proto.SellerResp) error {
	code, ok := c.Get(in.Email)
	if !ok || code != in.Code {
		out.Code = 401
		out.Msg = "Verify failed, try again"
		return nil
	}

	seller := dc.Seller{}
	dc.Db.Where("Email = ?", in.Email).First(&seller)
	md5Pwd := utils.Md5Sum(in.Password)
	if seller.Password != md5Pwd {
		out.Code = 401
		out.Msg = "Verify failed, try again"
		return nil
	}
	out.Code = 200
	out.Msg = "Welcome"
	out.Email = in.Email
}

func (SellerHandler) SellerLogout(ctx context.Context, in *proto.SellerLogoutRequest, out *proto.SellerResp) error {
	return nil
}
