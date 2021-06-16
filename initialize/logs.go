package initialize

import (
	"cronProject/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func GetLogger() *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   global.Config.Logs.FilePath,   // 日志文件路径
		MaxSize:    global.Config.Logs.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: global.Config.Logs.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     global.Config.Logs.MaxAge,     // 文件最多保存多少天
		Compress:   global.Config.Logs.Compress,   // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	var syncer zapcore.WriteSyncer
	var caller zap.Option
	var logger *zap.Logger
	if global.Config.Logs.Debug {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
		caller = zap.AddCaller()
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			syncer, atomicLevel,
		)
		development := zap.Development()
		filed := zap.Fields(zap.String("version", "1.0"))
		logger = zap.New(core, caller, development, filed)
	} else {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			syncer, atomicLevel,
		)
		development := zap.Development()
		filed := zap.Fields(zap.String("version", "1.0"))
		logger = zap.New(core, development, filed)
	}
	logger.Info("已启动")
	return logger
}
