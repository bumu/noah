package data

import (
	"noah/apps/infra/data/schema"

	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) {
	conn.Migrator().AutoMigrate(
		&schema.InfraLinuxMotd{},
		&schema.InfraHostHeartbeat{},
	)
}
