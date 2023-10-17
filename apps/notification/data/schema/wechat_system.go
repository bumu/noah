package schema

import "gorm.io/gorm"

type NotificationWechatSystem struct {
	gorm.Model

	Appid          string
	Secret         string
	Token          string
	EncodingAESKey string
	Admin          string
}

// TableName overrides the table name used by User to `profiles`
func (NotificationWechatSystem) TableName() string {
	return "tbl_notification_wechat_system"
}
