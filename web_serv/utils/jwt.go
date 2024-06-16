package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserToken struct {
	jwt.StandardClaims
	// 自定义用户信息
	UserName string `json:"user_name"`
}

// 客户token配置信息
var CustomerTokenExpireDuration = time.Hour
var CustomerSecretKey = []byte("customer_token")

// 前端用户token配置信息
var AdminTokenExpireDuration = time.Hour * 3
var AdminSecretKey = []byte("admin_token")

/**
 * @Description:  生成token
 * @param userName   用户名
 * @param userId     用户id
 * @param expireDuration   过期时间
 * @param key   加密的盐
 * @return string
 * @return error
 */
func GenToken(userName string, expireDuration time.Duration, secret_key []byte) (string, error) {

	user := UserToken{
		StandardClaims: jwt.StandardClaims{
			// 过期时间  现在的时间加上过期时间
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			Issuer:    "seckill_system",
		},
		UserName: userName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(secret_key)
}

/**
 * @Description:    认证token
 * @param tokenString    生成的token字符串
 * @param secret_key     加密的盐
 * @return *UserToken    返回一个结构体
 * @return error
 */
func VerifyToken(tokenString string, secret_key []byte) (*UserToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserToken{}, func(token *jwt.Token) (i interface{}, e error) {
		return secret_key, nil
	})

	if err != nil {
		return nil, err
	}

	userToken, is_ok := token.Claims.(*UserToken)
	// 验证token
	if token.Valid && is_ok {
		return userToken, nil
	}

	return nil, errors.New("token valid error!")

}
