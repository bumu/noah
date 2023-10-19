package data

import (
	"noah/apps/user/data/schema"

	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) {
	conn.Migrator().AutoMigrate(
		&schema.UserSession{},
	)
}
