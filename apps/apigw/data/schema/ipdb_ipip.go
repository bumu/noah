package schema

import "gorm.io/gorm"

// Ref: https://www.ip2location.com/database/db26-ip-country-region-city-latitude-longitude-zipcode-timezone-isp-domain-netspeed-areacode-weather-mobile-elevation-usagetype-addresstype-category-district-asn
type IpdbIpIp struct {
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
func (IpdbIpIp) TableName() string {
	return "tbl_ipdb_ipip"
}
