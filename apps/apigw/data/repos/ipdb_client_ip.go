package repos

import (
	"context"

	"apigw/apps/apigw/data/schema"

	"gorm.io/gorm"
)

type IpdbClientIpRepo struct {
	Conn *gorm.DB
}

func NewIpdbClientIpRepo(conn *gorm.DB) *IpdbClientIpRepo {
	return &IpdbClientIpRepo{conn}
}

func (r IpdbClientIpRepo) Create(ctx context.Context, entity *schema.IpdbClientIp) error {
	return r.Conn.Create(entity).Error
}

func (r IpdbClientIpRepo) First(ctx context.Context, token string) (*schema.IpdbClientIp, error) {
	dst := &schema.IpdbClientIp{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r IpdbClientIpRepo) List(ctx context.Context, limit, offset int) ([]*schema.IpdbClientIp, error) {
	dst := []*schema.IpdbClientIp{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
