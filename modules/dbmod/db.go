package dbmod

import (
	"apigw/pkg/configkit"

	"github.com/xo/dburl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConn() (*gorm.DB, error) {
	dsn, err := dburl.Parse(configkit.GlobalConfig.Database)
	if err != nil {
		return nil, err
	}

	return gorm.Open(mysql.Open(dsn.DSN), &gorm.Config{})
}
