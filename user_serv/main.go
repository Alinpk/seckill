package main

import (
	"user_serv/conf"
)

func ErrWapper(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	// init configuration
	conf.InitEnv()
	ErrWapper(conf.InitConfig())

	//
}