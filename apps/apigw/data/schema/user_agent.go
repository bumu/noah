package schema

import "gorm.io/gorm"

type Useragent struct {
	gorm.Model

	Useragent string
}

// TableName overrides the table name used by User to `profiles`
func (Useragent) TableName() string {
	return "tbl_useragent"
}
