package schema

import "gorm.io/gorm"

type Sensitive struct {
	gorm.Model

	Category    string
	CountryCode string // country code 3 letters
	Content     string `gorm:"unique:not null"`
	Remark      string
}

// TableName overrides the table name used by Key to `profiles`
func (Sensitive) TableName() string {
	return "tbl_sensitives"
}
