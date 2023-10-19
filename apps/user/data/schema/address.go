package schema

import "gorm.io/gorm"

type UserAddress struct {
	gorm.Model

	Coutry    string
	Province  string
	City      string
	Detail    string
	Unit      string
	BlackedAt int64 // 小黑屋
}

func (UserAddress) TableName() string {
	return "tbl_user_address"
}
