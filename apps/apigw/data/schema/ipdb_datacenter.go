package schema

import "gorm.io/gorm"

type IpdbDataCenter struct {
	gorm.Model

	Name            string
	Domain          string
	IPStart         string
	IPEnd           string
	created_ts      int
	last_updated_ts int
}

// TableName overrides the table name used by User to `profiles`
func (IpdbDataCenter) TableName() string {
	return "tbl_ipdb_datacenter"
}
