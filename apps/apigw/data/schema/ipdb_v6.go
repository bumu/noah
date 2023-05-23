package schema

import "gorm.io/gorm"

type IpdbV6 struct {
	gorm.Model

	IPRange            string
	Registry           string
	AssignedDate       int
	CountryCode2Letter string
}

// TableName overrides the table name used by User to `profiles`
func (IpdbV6) TableName() string {
	return "tbl_ipdb_v6"
}
