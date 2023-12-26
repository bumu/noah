package repos

import (
	"context"
	"time"

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

func (r InfraLinuxRepo) Count(ctx context.Context) (loginCnt int64, logoutCnt int64, cmdCnt int64, err error) {
	dst := []*schema.InfraLinuxMotd{}

	// loc, _ := time.LoadLocation("Asia/Singapore")
	// now := time.Now().In(loc)
	now := time.Now()

	// zero := now.Truncate(24 * time.Hour)

	// := now.Date().Format("2006-01-02")
	today := now.Truncate(24 * time.Hour).Add(-10 * time.Hour)
	tmr := now.Truncate(24 * time.Hour).Add(24 * time.Hour).Add(-10 * time.Hour)

	r.Conn.Debug().Where("action = ? and created_at BETWEEN ? AND ?", "login", today, tmr).
		Find(&dst).Count(&loginCnt)

	r.Conn.Debug().Where("action = ? and created_at BETWEEN ? AND ?", "logout", today, tmr).
		Find(&dst).Count(&logoutCnt)

	r.Conn.Debug().Where("created_at BETWEEN ? AND ?", "logout", today, tmr).Not("action = ?", "login").Not("action = ?", "logout").
		Find(&dst).Count(&cmdCnt)

	return loginCnt, logoutCnt, cmdCnt, err
}
