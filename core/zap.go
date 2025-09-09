package core

import (
	"fmt"
	"myblog-server-go/core/internal"
	"myblog-server-go/global"
	"myblog-server-go/util"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
func Zap() (log *zap.SugaredLogger) {
	if ok, _ := util.PathExists(global.BlogConfig.Zap.Director); !ok {
		_ = os.Mkdir(global.BlogConfig.Zap.Director, os.ModePerm)
		fmt.Printf("create %v directory\n", global.BlogConfig.Zap.Director)
	}
	levels := global.BlogConfig.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger := zap.New(zapcore.NewTee(cores...))
	if global.BlogConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(logger)
	return logger.Sugar()
}
