package schema

import "gorm.io/gorm"

// Ref: https://www.ip2location.com/database/ip2location
type IpdbIp2Location struct {
	gorm.Model

	IpFrom      uint32
	IpTo        uint32
	Country     string
	Region      string
	City        string
	Latitude    float32
	Longitude   float32
	Zipcode     string
	Timezone    string
	Domain      string
	NetSpeed    string
	AreaCode    string
	Weather     string
	Mobile      string
	Evlevation  string
	UsageType   string
	AddressType string
	Category    string
	District    string
	Asn         string
}

// TableName overrides the table name used by User to `profiles`
func (IpdbIp2Location) TableName() string {
	return "tbl_ipdb_ip2location"
}
