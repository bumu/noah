package repos

import (
	"context"

	"noah/apps/mall/data/schema"

	"gorm.io/gorm"
)

type MallShopRepo struct {
	Conn *gorm.DB
}

func NewMallShopRepo(conn *gorm.DB) *MallShopRepo {
	return &MallShopRepo{conn}
}

func (r MallShopRepo) Create(ctx context.Context, entity *schema.MallShop) error {
	return r.Conn.Create(entity).Error
}

func (r MallShopRepo) First(ctx context.Context, token string) (*schema.MallShop, error) {
	dst := &schema.MallShop{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r MallShopRepo) List(ctx context.Context, limit, offset int) ([]*schema.MallShop, error) {
	dst := []*schema.MallShop{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
