package schema

import "gorm.io/gorm"

type MallShop struct {
	gorm.Model

	LoginName string
	Password  string
}

// TableName overrides the table name used by User to `profiles`
func (MallShop) TableName() string {
	return "tbl_mall_shop"
}
