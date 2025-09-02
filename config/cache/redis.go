package cache

type Redis struct {
	Address  string `mapstructure:"address" json:"address" yaml:"address"`    // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // 数据库
}
