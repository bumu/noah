package schema

import "gorm.io/gorm"

// Singapore Standard Industrial Classification

type CompanySSICInfo struct {
	gorm.Model
	CompanySSICName     string `json:"ssic_name"`
	CompanySSICCode     string `json:"ssic_code"`
	CompanySSICClass    string `json:"ssic_class"`
	CompanySSICGroup    string `json:"ssic_group"`
	CompanySSICDivision string `json:"ssic_division"`
	CompanySSICSection  string `json:"ssic_section"`
}
