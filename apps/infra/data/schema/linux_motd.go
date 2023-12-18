package schema

import "gorm.io/gorm"

type InfraLinuxMotd struct {
	gorm.Model

	IP       string
	Hostname string
}

// TableName overrides the table name used by User to `profiles`
func (InfraLinuxMotd) TableName() string {
	return "tbl_infra_linux_motd"
}
