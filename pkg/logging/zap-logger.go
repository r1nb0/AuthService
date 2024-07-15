package logging

import (
	"github.com/fluxninja/lumberjack"
	"github.com/r1nb0/UserService/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type zapLogger struct {
	logger *zap.Logger
	cfg    *config.Config
}

func NewZapLogger(cfg *config.Config) Logger {
	logger := &zapLogger{
		cfg: cfg,
	}
	logger.Init()
	return logger
}

func (l *zapLogger) Init() {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.cfg.Logger.FilePath,
		MaxSize:    l.cfg.Logger.MaxSize,
		MaxBackups: l.cfg.Logger.MaxBackups,
		MaxAge:     l.cfg.Logger.MaxAge,
	})

	level := zap.NewAtomicLevelAt(zap.DebugLevel)
	prodCfg := zap.NewProductionEncoderConfig()
	prodCfg.TimeKey = "timestamp"
	prodCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	devCfg := zap.NewDevelopmentEncoderConfig()
	devCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(devCfg)
	fileEncoder := zapcore.NewJSONEncoder(prodCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)
	l.logger = zap.New(core)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, another map[string]interface{}) {
	fields := prepareLogInfo(cat, sub, another)
	l.logger.Info(msg, fields...)
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, another map[string]interface{}) {
	fields := prepareLogInfo(cat, sub, another)
	l.logger.Debug(msg, fields...)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, another map[string]interface{}) {
	fields := prepareLogInfo(cat, sub, another)
	l.logger.Warn(msg, fields...)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, another map[string]interface{}) {
	fields := prepareLogInfo(cat, sub, another)
	l.logger.Error(msg, fields...)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, another map[string]interface{}) {
	fields := prepareLogInfo(cat, sub, another)
	l.logger.Fatal(msg, fields...)
}

func prepareLogInfo(cat Category, sub SubCategory, another map[string]interface{}) []zap.Field {
	if another == nil {
		another = make(map[string]interface{})
	}
	another["Category"] = cat
	another["SubCategory"] = sub
	fields := make([]zap.Field, 0, len(another))
	for k, v := range another {
		fields = append(fields, zap.Any(k, v))
	}
	return fields
}
