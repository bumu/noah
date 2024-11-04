package repos

import (
	"context"
	"noah/apps/infra/data/schema"

	"gorm.io/gorm"
)

type InfraHostHeartbeatRepo struct {
	Conn *gorm.DB
}

func NewInfraHostHeartbeatRepo(conn *gorm.DB) *InfraHostHeartbeatRepo {
	return &InfraHostHeartbeatRepo{conn}
}

func (r InfraHostHeartbeatRepo) Create(ctx context.Context, entity *schema.InfraHostHeartbeat) error {
	return r.Conn.Create(entity).Error
}

func (r InfraHostHeartbeatRepo) Last(ctx context.Context) (*schema.InfraHostHeartbeat, error) {
	dst := &schema.InfraHostHeartbeat{}
	err := r.Conn.Last(dst).Error
	return dst, err
}

func (r InfraHostHeartbeatRepo) List(ctx context.Context, limit, offset int) ([]*schema.InfraHostHeartbeat, error) {
	dst := []*schema.InfraHostHeartbeat{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
