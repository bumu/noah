package schema

import "gorm.io/gorm"

type UserSession struct {
	gorm.Model

	Session    string
	Encryption string // AES
	IssuedAt   int64
	ExpiredAt  int64
}

// TableName overrides the table name used by User to `profiles`
func (UserSession) TableName() string {
	return "tbl_user_session"
}
