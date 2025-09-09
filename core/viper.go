package core

import (
	"flag"
	"fmt"
	"myblog-server-go/core/internal"
	"myblog-server-go/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper
//
//	@Description: 初始化 viper 配置，读取单一配置文件
//	@return *viper.Viper
func Viper() *viper.Viper {
	config := getConfigPath()

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.BlogConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.BlogConfig); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}

	return v
}

// getConfigPath
//
//	@Description: 获取配置文件路径，优先级: 命令行 > 环境变量 > 默认值。其中，命令行参数 `-p` 用于开启发布模式，不需要传值。
//	@return config
func getConfigPath() (config string) {
	// `-c` flag parse
	release := "null"
	flag.StringVar(&release, "p", "", "choose config file.")
	flag.Parse()
	if release == "release" { // 发布模式
		config = internal.ConfigReleaseFile
		fmt.Printf("正在使用命令行的 '-p' 参数传递的值, config 的路径为 %s\n", config)
		return
	}

	if env := os.Getenv(internal.ConfigEnv); env != "" { // 判断环境变量 GVA_CONFIG
		config = env
		fmt.Printf("正在使用 %s 环境变量, config 的路径为 %s\n", internal.ConfigEnv, config)
		return
	}

	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		config = internal.ConfigDefaultFile
		fmt.Printf("配置文件路径不存在, 使用默认配置文件路径: %s\n", config)
	}

	return
}
