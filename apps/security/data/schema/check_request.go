package schema

import (
	"gorm.io/gorm"
)

// Device is used to store device information
type SecurityCheckRequest struct {
	gorm.Model

	RequestID      string
	RemoteAddr     string
	UserAgent      string
	FingerprintJa3 string
	FingerprintH2  string
	Protocol       string
	Method         string
	CheckURL       string
	CheckHeader    string
	CheckBody      string
}

func (SecurityCheckRequest) TableName() string {
	return "tbl_security_check_request"
}
