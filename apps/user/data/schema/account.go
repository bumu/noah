package schema

import "gorm.io/gorm"

type UserAccount struct {
	gorm.Model

	OpenID    string
	Nickname  string
	BlackedAt int64 // 小黑屋
}

func (UserAccount) TableName() string {
	return "tbl_user_account"
}
