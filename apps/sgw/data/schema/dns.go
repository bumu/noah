package schema

import (
	"gorm.io/gorm"
)

type GatewayDNSRecord struct {
	gorm.Model

	Domain       string
	DomainPrefix string
	RecordType   string
	RecordValue  string
	TTL          int64
	Comment      string
}

// TableName overrides the table name used by User to `profiles`
func (GatewayDNSRecord) TableName() string {
	return "tbl_gateway_record"
}
