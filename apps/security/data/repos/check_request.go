package repos

import (
	"context"

	"noah/apps/security/data/schema"

	"gorm.io/gorm"
)

type SecurityCheckRequestRepo struct {
	Conn *gorm.DB
}

func NewSecurityCheckRequestRepo(conn *gorm.DB) *SecurityCheckRequestRepo {
	return &SecurityCheckRequestRepo{conn}
}

func (r SecurityCheckRequestRepo) Create(ctx context.Context, entity *schema.SecurityCheckRequest) error {
	return r.Conn.Create(entity).Error
}

func (r SecurityCheckRequestRepo) First(ctx context.Context, token string) (*schema.SecurityCheckRequest, error) {
	dst := &schema.SecurityCheckRequest{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r SecurityCheckRequestRepo) Last(ctx context.Context) (*schema.SecurityCheckRequest, error) {
	dst := &schema.SecurityCheckRequest{}
	err := r.Conn.Last(dst).Error
	return dst, err
}
