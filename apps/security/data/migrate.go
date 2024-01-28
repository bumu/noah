package data

import (
	"noah/apps/security/data/schema"

	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) {
	conn.Migrator().AutoMigrate(
		&schema.SecurityDevice{},
		&schema.SecurityUserAgent{},
		&schema.SecurityCheckRequest{},
	)
}
