package repos

import (
	"context"

	"apigw/apps/apigw/data/schema"

	"gorm.io/gorm"
)

type SensitiveRepo struct {
	Conn *gorm.DB
}

func NewSensitiveRepo(conn *gorm.DB) *SensitiveRepo {
	return &SensitiveRepo{conn}
}

func (r SensitiveRepo) Create(ctx context.Context, entity *schema.Sensitive) error {
	return r.Conn.Create(entity).Error
}

func (r SensitiveRepo) First(ctx context.Context, token string) (*schema.Sensitive, error) {
	dst := &schema.Sensitive{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r SensitiveRepo) List(ctx context.Context, limit, offset int) ([]*schema.Sensitive, error) {
	dst := []*schema.Sensitive{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
