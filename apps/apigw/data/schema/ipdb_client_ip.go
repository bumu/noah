package schema

import "gorm.io/gorm"

type IpdbClientIp struct {
	gorm.Model

	Ip           string
	IspDomain    string
	IspName      string
	IpCity       string
	Cnt          int64
	IsDataCenter bool
}

// TableName overrides the table name used by User to `profiles`
func (IpdbClientIp) TableName() string {
	return "tbl_ipdb_client_ip"
}
