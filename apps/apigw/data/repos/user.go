package repos

import (
	"context"

	"apigw/apps/apigw/data/schema"

	"gorm.io/gorm"
)

type UserRepo struct {
	Conn *gorm.DB
}

func NewUserRepo(conn *gorm.DB) *UserRepo {
	return &UserRepo{conn}
}

func (r UserRepo) Create(ctx context.Context, entity *schema.User) error {
	return r.Conn.Create(entity).Error
}

func (r UserRepo) List(ctx context.Context, limit, offset int) ([]*schema.User, error) {
	dst := []*schema.User{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
