package repos

import (
	"context"

	"noah/apps/security/data/schema"

	"gorm.io/gorm"
)

type SecurityUserAgentRepo struct {
	Conn *gorm.DB
}

func NewSecurityUserAgentRepo(conn *gorm.DB) *SecurityUserAgentRepo {
	return &SecurityUserAgentRepo{conn}
}

func (r SecurityUserAgentRepo) Create(ctx context.Context, entity *schema.SecurityUserAgent) error {
	return r.Conn.Create(entity).Error
}

func (r SecurityUserAgentRepo) First(ctx context.Context, token string) (*schema.SecurityUserAgent, error) {
	dst := &schema.SecurityUserAgent{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r SecurityUserAgentRepo) Last(ctx context.Context) (*schema.SecurityUserAgent, error) {
	dst := &schema.SecurityUserAgent{}
	err := r.Conn.Last(dst).Error
	return dst, err
}
