package schema

import "gorm.io/gorm"

type InfraHostHeartbeat struct {
	gorm.Model

	Hostname    string
	IP          string
	Arch        string
	OS          string
	Kernel      string
	SystemInfo  string
	Environment string
}

// TableName overrides the table name used by InfraHostHeartbeat
func (InfraHostHeartbeat) TableName() string {
	return "tbl_infra_host_heartbeat"
}
