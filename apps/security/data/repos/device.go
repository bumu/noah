package repos

import (
	"context"

	"noah/apps/security/data/schema"

	"gorm.io/gorm"
)

type SecurityDeviceRepo struct {
	Conn *gorm.DB
}

func NewSecurityDeviceRepo(conn *gorm.DB) *SecurityDeviceRepo {
	return &SecurityDeviceRepo{conn}
}

func (r SecurityDeviceRepo) Create(ctx context.Context, entity *schema.UserSession) error {
	return r.Conn.Create(entity).Error
}

func (r SecurityDeviceRepo) First(ctx context.Context, token string) (*schema.UserSession, error) {
	dst := &schema.UserSession{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r SecurityDeviceRepo) List(ctx context.Context, limit, offset int) ([]*schema.UserSession, error) {
	dst := []*schema.UserSession{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
