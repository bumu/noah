package dbmod

import (
	"apigw/pkg/configkit"
	"fmt"

	"github.com/xo/dburl"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConn() (*gorm.DB, error) {
	dsn, err := dburl.Parse(configkit.GlobalConfig.DatabaseURL)
	if err != nil {
		return nil, err
	}

	switch dsn.Driver {
	case "postgres":
		return gorm.Open(postgres.Open(dsn.DSN), &gorm.Config{})
	case "mysql":
		return gorm.Open(mysql.Open(dsn.DSN), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unknown driver: %s", dsn.Driver)
	}

}
