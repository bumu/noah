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
