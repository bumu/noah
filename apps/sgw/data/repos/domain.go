package repos

import (
	"context"

	"noah/apps/sgw/data/schema"

	"gorm.io/gorm"
)

type GatewayDomainRepo struct {
	Conn *gorm.DB
}

func NewGatewayDomainRepo(conn *gorm.DB) *GatewayDomainRepo {
	return &GatewayDomainRepo{conn}
}

func (r GatewayDomainRepo) Create(ctx context.Context, entity *schema.GatewayDomain) error {
	return r.Conn.Create(entity).Error
}

func (r GatewayDomainRepo) First(ctx context.Context, token string) (*schema.GatewayDomain, error) {
	dst := &schema.GatewayDomain{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r GatewayDomainRepo) List(ctx context.Context, limit, offset int) ([]*schema.GatewayDomain, error) {
	dst := []*schema.GatewayDomain{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
