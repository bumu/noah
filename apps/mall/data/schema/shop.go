package schema

import "gorm.io/gorm"

type UserSession struct {
	gorm.Model

	LoginName string
	Password  string
}

// TableName overrides the table name used by User to `profiles`
func (UserSession) TableName() string {
	return "tbl_mall_shop"
}
