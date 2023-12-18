package repos

import (
	"context"

	"noah/apps/infra/data/schema"

	"gorm.io/gorm"
)

type InfraLinuxRepo struct {
	Conn *gorm.DB
}

func NewInfraLinuxRepo(conn *gorm.DB) *InfraLinuxRepo {
	return &InfraLinuxRepo{conn}
}

func (r InfraLinuxRepo) Create(ctx context.Context, entity *schema.InfraLinuxMotd) error {
	return r.Conn.Create(entity).Error
}

func (r InfraLinuxRepo) First(ctx context.Context, token string) (*schema.InfraLinuxMotd, error) {
	dst := &schema.InfraLinuxMotd{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r InfraLinuxRepo) List(ctx context.Context, limit, offset int) ([]*schema.InfraLinuxMotd, error) {
	dst := []*schema.InfraLinuxMotd{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
