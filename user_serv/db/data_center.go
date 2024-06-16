package dc

import (
	"user_serv/conf"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func InitDataCenter() error {
	sqlConf := conf.GetSqlConf()
	var err error
	Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		sqlConf.User,
		sqlConf.Pwd,
		sqlConf.Addr,
		sqlConf.DbName,
	))

	Db.LogMode(sqlConf.LogMode)

	Db.DB().SetMaxOpenConns(100) // 最大连接数
	Db.DB().SetMaxIdleConns(50)  // 最大空闲数

	return err
}