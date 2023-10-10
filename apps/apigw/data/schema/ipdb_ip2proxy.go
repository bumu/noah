package schema

import "gorm.io/gorm"

// Ref: https://www.ip2location.com/database/ip2proxy
type IpdbIp2Proxy struct {
	gorm.Model

	Ip          string
	ProxyType   string
	Country     string
	Region      string
	City        string
	Isp         string
	Domain      string
	UsageType   string
	Asn         string
	LastSeen    string
	Threat      string
	Residential string
	Provider    string
}

// TableName overrides the table name used by User to `profiles`
func (IpdbIp2Proxy) TableName() string {
	return "tbl_ipdb_ip2proxy"
}
