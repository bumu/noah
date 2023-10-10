package schema

import "gorm.io/gorm"

// Refer: https://www.ip2location.com/vpn-provider-coverage
type IpdbCloudProvider struct {
	gorm.Model

	ServiceType string
	Name        string
	Website     string
}

// TableName overrides the table name used by User to `profiles`
func (IpdbCloudProvider) TableName() string {
	return "tbl_ipdb_cloud_provider"
}
