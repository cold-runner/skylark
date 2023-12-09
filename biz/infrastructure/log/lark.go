package log

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/biz/config"

	hertzzap "github.com/hertz-contrib/logger/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

func Init(c config.Conf) {
	// developMode config
	loggerLark := hertzzap.NewLogger(
		hertzzap.WithCores(developCoreConfig(c.GetLog())...),
		hertzzap.WithZapOptions(
			zap.AddCallerSkip(3),
			zap.WithCaller(true),
			// zap.AddStacktrace(zap.DebugLevel), // 打印堆栈调用
		),
	)

	hlog.SetLogger(loggerLark)
}

// TODO 配置自检
func developCoreConfig(c *config.Log) []hertzzap.CoreConfig {
	// log level
	level, _ := zap.ParseAtomicLevel(c.Level)
	lvl := zap.NewAtomicLevelAt(level.Level())

	// encoder config
	var enc zapcore.Encoder
	switch c.Format {
	case "console":
		enc = zapcore.NewConsoleEncoder(encoderConfig())
	case "json":
		enc = zapcore.NewJSONEncoder(encoderConfig())
	default:
		panic("不支持的日志格式")
	}

	// output path
	paths := strings.Split(c.OutputPaths, ",")[1:]
	outPaths := make([]zapcore.WriteSyncer, len(paths)+1)
	outPaths[0] = zapcore.AddSync(os.Stdout)
	for i := 1; i < len(outPaths); i++ {
		filePath := "log/" + paths[i-1]
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		outPaths[i] = zapcore.AddSync(file)
	}

	zapCoreConfigs := make([]hertzzap.CoreConfig, len(outPaths))
	for i := range zapCoreConfigs {
		zapCoreConfigs[i] = hertzzap.CoreConfig{
			Enc: enc,
			Ws:  outPaths[i],
			Lvl: lvl,
		}
	}

	return zapCoreConfigs
}

// encoderConfig encoder config for debug
func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "name",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
