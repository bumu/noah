package repos

import (
	"context"

	"noah/apps/mall/data/schema"

	"gorm.io/gorm"
)

type UserSessionRepo struct {
	Conn *gorm.DB
}

func NewUserSessionRepo(conn *gorm.DB) *UserSessionRepo {
	return &UserSessionRepo{conn}
}

func (r UserSessionRepo) Create(ctx context.Context, entity *schema.UserSession) error {
	return r.Conn.Create(entity).Error
}

func (r UserSessionRepo) First(ctx context.Context, token string) (*schema.UserSession, error) {
	dst := &schema.UserSession{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r UserSessionRepo) List(ctx context.Context, limit, offset int) ([]*schema.UserSession, error) {
	dst := []*schema.UserSession{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
