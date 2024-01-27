package schema

import (
	"gorm.io/gorm"
)

// Device is used to store device information
type SecurityUserAgent struct {
	gorm.Model

	UAType           string
	UA               string
	UADevice         string
	UAOS             string
	UAClassification string
	UABrowserName    string
	FingerprintJa3   string
	FingerprintH2    string
}

func (SecurityUserAgent) TableName() string {
	return "tbl_security_user_agent"
}
