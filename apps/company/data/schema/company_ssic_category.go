package schema

import "gorm.io/gorm"

// Singapore Standard Industrial Classification

type CompanySSICCategory struct {
	gorm.Model
	CompanySSICClassification string `json:"ssic_classification"`
	CompanySSICCategoryName   string `json:"ssic_category_name"`
	CompanySSICCategoryCode   string `json:"ssic_category_code"`
}
