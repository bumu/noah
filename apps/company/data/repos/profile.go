package repos

import (
	"context"

	"noah/apps/company/data/schema"

	"gorm.io/gorm"
)

type CompanyProfileRepo struct {
	Conn *gorm.DB
}

func NewCompanyProfileRepo(conn *gorm.DB) *CompanyProfileRepo {
	return &CompanyProfileRepo{conn}
}

func (r CompanyProfileRepo) Create(ctx context.Context, entity *schema.CompnayProfile) error {
	return r.Conn.Create(entity).Error
}

func (r CompanyProfileRepo) First(ctx context.Context, token string) (*schema.CompnayProfile, error) {
	dst := &schema.CompnayProfile{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r CompanyProfileRepo) List(ctx context.Context, limit, offset int) ([]*schema.CompnayProfile, error) {
	dst := []*schema.CompnayProfile{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
