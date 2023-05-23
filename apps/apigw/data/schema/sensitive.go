package schema

import "gorm.io/gorm"

type Sensitive struct {
	gorm.Model

	Title  string
	Vendor string // openai azure
	Token  string
}

// TableName overrides the table name used by Key to `profiles`
func (Sensitive) TableName() string {
	return "tbl_sensitives"
}
