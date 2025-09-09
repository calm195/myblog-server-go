package internal

import (
	"myblog-server-go/config/orm"
	"myblog-server-go/global"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config
//
//	@Description:
//	@receiver g
//	@param prefix
//	@param singular
//	@return *orm.Config
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	var general orm.GeneralDB
	switch global.BlogConfig.System.DbType {
	case "mysql":
		general = global.BlogConfig.Mysql.GeneralDB
	case "pgsql":
		general = global.BlogConfig.Pgsql.GeneralDB
	default:
		general = global.BlogConfig.Mysql.GeneralDB
	}
	return &gorm.Config{
		Logger: logger.New(NewWriter(general), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      general.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
