package dbmod

import (
	"fmt"

	"noah/pkg/configkit"

	"github.com/xo/dburl"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConn() (*gorm.DB, error) {
	dialector, err := databaseDialector(configkit.GlobalConfig.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return gorm.Open(dialector, &gorm.Config{})
}

func databaseDialector(databaseURL string) (gorm.Dialector, error) {
	dsn, err := dburl.Parse(databaseURL)
	if err != nil {
		return nil, err
	}

	switch dsn.Dialect {
	case "postgres":
		return postgres.Open(dsn.DSN), nil
	case "mysql":
		return mysql.Open(dsn.DSN), nil
	default:
		return nil, fmt.Errorf("unsupported database dialect: %s", dsn.Dialect)
	}

}
