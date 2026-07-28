package repos

import (
	"context"

	"noah/apps/company/data/schema"

	"gorm.io/gorm"
)

type CompanyUploadRepo struct {
	conn *gorm.DB
}

func NewCompanyUploadRepo(conn *gorm.DB) *CompanyUploadRepo {
	return &CompanyUploadRepo{conn: conn}
}

func (r CompanyUploadRepo) Create(ctx context.Context, upload *schema.CompanyUpload) error {
	return r.conn.WithContext(ctx).Create(upload).Error
}
