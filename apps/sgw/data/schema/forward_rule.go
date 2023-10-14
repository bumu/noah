package schema

import "gorm.io/gorm"

type GatewayForward struct {
	gorm.Model

	Host        string
	Port        int64
	ForwardType string
	Path        string
	Upstream    string
	Comment     string
}

// TableName overrides the table name used by User to `profiles`
func (GatewayForward) TableName() string {
	return "tbl_gateway_forward"
}
