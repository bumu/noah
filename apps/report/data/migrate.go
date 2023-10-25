package data

import (
	"noah/apps/apigw/data/schema"

	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) {
	conn.Migrator().AutoMigrate(
		&schema.Key{},
		&schema.User{},
		&schema.Ipdb{},
		&schema.IpdbV6{},
		&schema.Sensitive{},
		&schema.Useragent{},
		&schema.UseragentOS{},
		&schema.IpdbDataCenter{},
		&schema.IpdbClientIp{},
		&schema.IpdbCloudProvider{},
		&schema.IpdbIpIp{},
		&schema.IpdbIp2Proxy{},
		&schema.IpdbIp2Location{},
	)
}
