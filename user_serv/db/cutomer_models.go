package dc

import "time"

/*
用户模型类
*/
type Customer struct {
	Id         int    `gorm:"type:int(10);AUTO_INCREMENT"`
	Email      string `gorm:"type:varchar(64)"`
	Password   string `gorm:"type:varchar(64)"`
	Status     int    `gorm:"type:int(2)"`
	CreateTime time.Time
}

// 指定表名
func (Customer) TableName() string {
	return "customer"
}
