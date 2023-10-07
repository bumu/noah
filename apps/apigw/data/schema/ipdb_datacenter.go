package schema

import "gorm.io/gorm"

type IpdbDataCenter struct {
	gorm.Model

	Name      string
	Domain    string
	IpStart   string
	IpEnd     string
	CreatedTs int64
	UpdatedTs int64
}

// TableName overrides the table name used by User to `profiles`
func (IpdbDataCenter) TableName() string {
	return "tbl_ipdb_datacenter"
}
