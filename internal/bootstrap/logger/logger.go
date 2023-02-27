package logger

import (
	"fmt"
	microconfig "go-micro.dev/v4/config"
	"go-micro.dev/v4/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"

	zaplogger "chat/internal/bootstrap/logger/zap"
	"chat/internal/config"
)

var Module = fx.Options(
	fx.Provide(configureLogger),
	fx.WithLogger(provideContainerLogger),
	fx.Provide(NewLogger),
)

type loggerConf struct {
	Level  LogLevel  `json:"level,omitempty"`
	Output LogOutput `json:"output,omitempty"`
	Path   string    `json:"path,omitempty"`
	Backup int       `json:"backup,omitempty"`
	Size   int       `json:"size,omitempty"`
	Age    int       `json:"age,omitempty"`
}

func NewLogger(globalLogger *zap.Logger) (logger.Logger, error) {
	log, err := zaplogger.NewLogger(zaplogger.WithZap(globalLogger))
	if err != nil {
		return nil, err
	}
	logger.DefaultLogger = log
	return log, nil
}

func provideContainerLogger() fxevent.Logger {
	return &fxevent.NopLogger
}

func configureLogger(info *config.ServiceInfo, loader microconfig.Config) (*zap.Logger, error) {
	var conf loggerConf
	if err := loader.Get("logger").Scan(&conf); err != nil {
		return nil, err
	}
	enc := zapcore.NewJSONEncoder(encoderConfig())
	if info.Development() {
		enc = zapcore.NewConsoleEncoder(encoderConfig())
	}
	w := zapcore.AddSync(os.Stdout)
	if conf.Output == LogOutputFile {
		w = zapcore.NewMultiWriteSyncer(w, zapcore.AddSync(filterWriter(info, &conf)))
	}
	core := zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), level(conf.Level, zapcore.InfoLevel))
	global := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(
			zap.String("service.name", info.Name),
			zap.String("service.id", info.Id),
			zap.String("service.version", info.Version),
		),
	)
	zap.ReplaceGlobals(global)
	return global, nil
}

func encoderConfig() zapcore.EncoderConfig {
	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.TimeKey = "timestamp"
	return conf
}

func level(level LogLevel, defaultVal zapcore.Level) zapcore.Level {
	switch level {
	case LogLevelDebug:
		return zapcore.DebugLevel
	case LogLevelInfo:
		return zapcore.InfoLevel
	case LogLevelWarn:
		return zapcore.WarnLevel
	case LogLevelError:
		return zapcore.ErrorLevel
	default:
		return defaultVal
	}
}

func filterWriter(info *config.ServiceInfo, conf *loggerConf) io.Writer {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", conf.Path, info.Name),
		MaxSize:    conf.Size,
		MaxAge:     conf.Age,
		MaxBackups: conf.Backup,
		Compress:   false,
	}
}
