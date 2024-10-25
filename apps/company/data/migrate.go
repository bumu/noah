package data

import (
	"noah/apps/company/data/schema"

	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) {
	conn.Migrator().AutoMigrate(
		schema.CompnayProfile{},
	)
}
