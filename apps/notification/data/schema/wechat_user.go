package schema

import "gorm.io/gorm"

type NotificationWechatUser struct {
	gorm.Model

	Openid            string
	Unionid           string
	Nickname          string
	wxid              string
	NotificationGroup string
}

// TableName overrides the table name used by User to `profiles`
func (NotificationWechatUser) TableName() string {
	return "tbl_notification_wechat_user"
}
