package data

import (
	"noah/apps/sgw/data/schema"

	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) {
	conn.Migrator().AutoMigrate(
		&schema.GatewayCert{},
		&schema.GatewayDomain{},
		&schema.GatewayDNSRecord{},
		&schema.GatewayForward{},
	)
}
