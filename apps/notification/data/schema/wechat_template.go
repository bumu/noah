package schema

import "gorm.io/gorm"

type NotificationWechatTemplate struct {
	gorm.Model

	TemplateId      string
	TemplateTitle   string
	TemplateContent string
	TimeSpec        string
}

// TableName overrides the table name used by User to `profiles`
func (NotificationWechatTemplate) TableName() string {
	return "tbl_notification_wechat_template"
}
