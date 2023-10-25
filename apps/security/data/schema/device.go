package schema

import "gorm.io/gorm"

type SecurityDevice struct {
	gorm.Model

	DeviceType        string
	DeviceName        string
	DeviceBrand       string
	DeviceModel       string
	DeviceIMEI        string
	DeviceSerial      string
	DeviceMacAddress  string
	DeviceOS          string
	DeviceOSVersion   string
	DeviceFingerprint string
}

// TableName overrides the table name used by User to `profiles`
func (SecurityDevice) TableName() string {
	return "tbl_security_device"
}
