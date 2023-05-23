package data

import (
	"apigw/apps/apigw/data/schema"

	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) {
	conn.Migrator().AutoMigrate(
		&schema.Key{},
		&schema.User{},
		&schema.Ipdb{},
		&schema.IpdbV6{},
		&schema.Sensitive{},
	)
}
