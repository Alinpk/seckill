package dc

import "time"

type Seller struct {
	Id         int    `gorm:"type:int(10);AUTO_INCREMENT"`
	Email      string `gorm:"type:varchar(64)"`
	Password   string `gorm:"type:varchar(64)"`
	Status     int    `gorm:"type:int(2)"`
	CreateTime time.Time
}

/*
me@private.com
123456
*/

func (Seller) TableName() string {
	return "Seller"
}
