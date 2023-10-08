package repos

import (
	"context"

	"apigw/apps/apigw/data/schema"

	"gorm.io/gorm"
)

type UseragentRepo struct {
	Conn *gorm.DB
}

func NewUseragentRepo(conn *gorm.DB) *UseragentRepo {
	return &UseragentRepo{conn}
}

func (r UseragentRepo) Create(ctx context.Context, entity *schema.Useragent) error {
	return r.Conn.Create(entity).Error
}

func (r UseragentRepo) First(ctx context.Context, token string) (*schema.Useragent, error) {
	dst := &schema.Useragent{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r UseragentRepo) List(ctx context.Context, limit, offset int) ([]*schema.Useragent, error) {
	dst := []*schema.Useragent{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}

func (r UseragentRepo) GetDataCenter(ctx context.Context, ip uint32) (*schema.Useragent, error) {
	dst := &schema.Useragent{}
	err := r.Conn.Where("ip_start_int <= ? and ip_end_int >= ?", ip, ip).First(&dst).Error
	return dst, err
}
