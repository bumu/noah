package repos

import (
	"context"

	"noah/apps/notification/data/schema"

	"gorm.io/gorm"
)

type NotificationWechatUserRepo struct {
	Conn *gorm.DB
}

func NewNotificationWechatUserRepo(conn *gorm.DB) *NotificationWechatUserRepo {
	return &NotificationWechatUserRepo{conn}
}

func (r NotificationWechatUserRepo) Create(ctx context.Context, entity *schema.NotificationWechatUser) error {
	return r.Conn.Create(entity).Error
}

func (r NotificationWechatUserRepo) First(ctx context.Context, token string) (*schema.NotificationWechatUser, error) {
	dst := &schema.NotificationWechatUser{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r NotificationWechatUserRepo) List(ctx context.Context, limit, offset int) ([]*schema.NotificationWechatUser, error) {
	dst := []*schema.NotificationWechatUser{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
