package schema

import "gorm.io/gorm"

type UseragentOS struct {
	gorm.Model

	Useragent string
	OS        string
}

// TableName overrides the table name used by User to `profiles`
func (UseragentOS) TableName() string {
	return "tbl_useragent_os"
}
