package schema

import (
	"time"

	"gorm.io/gorm"
)

type GatewayDomain struct {
	gorm.Model

	Domain    string
	Issuer    string
	Expired   bool
	ExpiredAt time.Time
}

// TableName overrides the table name used by User to `profiles`
func (GatewayDomain) TableName() string {
	return "tbl_gateway_domain"
}
