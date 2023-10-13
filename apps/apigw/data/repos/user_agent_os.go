package repos

import (
	"context"

	"noah/apps/apigw/data/schema"

	"gorm.io/gorm"
)

type UseragentOSRepo struct {
	Conn *gorm.DB
}

func NewUseragentOSRepo(conn *gorm.DB) *UseragentOSRepo {
	return &UseragentOSRepo{conn}
}

func (r UseragentOSRepo) Create(ctx context.Context, entity *schema.UseragentOS) error {
	return r.Conn.Create(entity).Error
}

func (r UseragentOSRepo) First(ctx context.Context, token string) (*schema.UseragentOS, error) {
	dst := &schema.UseragentOS{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r UseragentOSRepo) List(ctx context.Context, limit, offset int) ([]*schema.UseragentOS, error) {
	dst := []*schema.UseragentOS{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}

func (r UseragentOSRepo) GetDataCenter(ctx context.Context, ip uint32) (*schema.UseragentOS, error) {
	dst := &schema.UseragentOS{}
	err := r.Conn.Where("ip_start_int <= ? and ip_end_int >= ?", ip, ip).First(&dst).Error
	return dst, err
}
