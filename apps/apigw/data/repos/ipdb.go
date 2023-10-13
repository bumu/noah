package repos

import (
	"context"

	"noah/apps/apigw/data/schema"

	"gorm.io/gorm"
)

type IpdbRepo struct {
	Conn *gorm.DB
}

func NewIpdbRepo(conn *gorm.DB) *IpdbRepo {
	return &IpdbRepo{conn}
}

func (r IpdbRepo) Create(ctx context.Context, entity *schema.Key) error {
	return r.Conn.Create(entity).Error
}

func (r IpdbRepo) First(ctx context.Context, token string) (*schema.Ipdb, error) {
	dst := &schema.Ipdb{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r IpdbRepo) List(ctx context.Context, limit, offset int) ([]*schema.Ipdb, error) {
	dst := []*schema.Ipdb{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
