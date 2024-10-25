package schema

import "gorm.io/gorm"

type CompnayProfile struct {
	gorm.Model

	Name    string `json:"name" gorm:"unique"` // Name
	UEN     string `json:"uen"`                // UEN
	Address string `json:"address"`            // 详细地址
	Website string `json:"website"`            // 官网
	Phone   string `json:"phone"`              // 电话
	Status  int32  `json:"status"`             // Status
}

// TableName overrides the table name used by User to `profiles`
func (CompnayProfile) TableName() string {
	return "tbl_company_profile"
}
