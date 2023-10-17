package data

import (
	"noah/apps/notification/data/schema"

	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) {
	conn.Migrator().AutoMigrate(
		&schema.NotificationWechatSystem{},
		&schema.NotificationWechatUser{},
		&schema.NotificationWechatTemplate{},
	)
}
