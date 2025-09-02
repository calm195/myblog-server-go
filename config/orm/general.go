package orm

import (
	"strings"

	"gorm.io/gorm/logger"
)

type DsnProvider interface {
	Dsn() string
}

// GeneralDB 也被 Pgsql 和 Mysql 原样使用
type GeneralDB struct {
	Path               string `mapstructure:"path" json:"path" yaml:"path"`                            // 数据库主机地址
	Port               string `mapstructure:"port" json:"port" yaml:"port"`                            // 数据库端口
	Username           string `mapstructure:"username" json:"username" yaml:"username"`                // 登录用户名
	Password           string `mapstructure:"password" json:"password" yaml:"password"`                // 登录密码
	Dbname             string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                   // 数据库名
	Config             string `mapstructure:"config" json:"config" yaml:"config"`                      // 连接参数
	Prefix             string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                      // 数据库前缀
	Engine             string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`     // 数据库引擎，默认InnoDB
	MaxIdleConnections int    `mapstructure:"max-idle" json:"max-idle" yaml:"max-idle"`                // 空闲中的最大连接数
	MaxOpenConnections int    `mapstructure:"max-open" json:"max-open" yaml:"max-open"`                // 打开到数据库的最大连接数
	Singular           bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                // 是否开启全局禁用复数，true表示开启
	LogMode            string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode" default:"info"` // 是否开启Gorm全局日志
	LogZap             bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                   // 是否通过zap写入日志文件
}

func (c GeneralDB) LogLevel() logger.LogLevel {
	switch strings.ToLower(c.LogMode) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}
