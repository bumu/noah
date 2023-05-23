package schema

import "gorm.io/gorm"

type Ipdb struct {
	gorm.Model

	IPFrom             int
	IPTo               int
	Registry           string
	AssignedDate       int
	CountryCode2Letter string
	CountryCode3Letter string
	Country            string
}

// TableName overrides the table name used by User to `profiles`
func (Ipdb) TableName() string {
	return "tbl_ipdb"
}
