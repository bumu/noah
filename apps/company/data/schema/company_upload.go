package schema

import (
	"time"

	"gorm.io/datatypes"
)

type CompanyUpload struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Source    string         `json:"source" gorm:"size:64;not null;index"`
	Data      datatypes.JSON `json:"data" gorm:"type:json;not null"`
	CreatedAt time.Time      `json:"createdAt"`
}

func (CompanyUpload) TableName() string {
	return "company_upload"
}
