package config

import (
	"myblog-server-go/config/cache"
	"myblog-server-go/config/orm"
	"myblog-server-go/config/oss"
)

type Server struct {
	Zap     Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   cache.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mongo   cache.Mongo `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	System  System      `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha     `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// orm
	Mysql orm.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql orm.Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	// oss
	Local     oss.Local     `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu     oss.Qiniu     `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS oss.AliyunOSS `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`

	DiskList []DiskList `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
