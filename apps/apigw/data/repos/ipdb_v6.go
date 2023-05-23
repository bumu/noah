package repos

import (
	"context"

	"apigw/apps/apigw/data/schema"

	"gorm.io/gorm"
)

type IpdbV6Repo struct {
	Conn *gorm.DB
}

func NewIpdbV6Repo(conn *gorm.DB) *IpdbV6Repo {
	return &IpdbV6Repo{conn}
}

func (r IpdbV6Repo) Create(ctx context.Context, entity *schema.Key) error {
	return r.Conn.Create(entity).Error
}

func (r IpdbV6Repo) First(ctx context.Context, token string) (*schema.IpdbV6, error) {
	dst := &schema.IpdbV6{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r IpdbV6Repo) List(ctx context.Context, limit, offset int) ([]*schema.IpdbV6, error) {
	dst := []*schema.IpdbV6{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
