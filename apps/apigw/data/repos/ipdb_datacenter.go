package repos

import (
	"context"

	"apigw/apps/apigw/data/schema"

	"gorm.io/gorm"
)

type IpdbDataCenterRepo struct {
	Conn *gorm.DB
}

func NewIpdbDataCenterRepo(conn *gorm.DB) *IpdbDataCenterRepo {
	return &IpdbDataCenterRepo{conn}
}

func (r IpdbDataCenterRepo) Create(ctx context.Context, entity *schema.Key) error {
	return r.Conn.Create(entity).Error
}

func (r IpdbDataCenterRepo) First(ctx context.Context, token string) (*schema.IpdbDataCenter, error) {
	dst := &schema.IpdbDataCenter{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r IpdbDataCenterRepo) List(ctx context.Context, limit, offset int) ([]*schema.IpdbDataCenter, error) {
	dst := []*schema.IpdbDataCenter{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
