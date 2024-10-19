package logging

import (
	"github.com/erfanwd/golang-course-project/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapSinLogger *zap.SugaredLogger

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

var logLevelMap = map[string]zapcore.Level{
	"debug":   zapcore.DebugLevel,
	"info":    zapcore.InfoLevel,
	"warning": zapcore.WarnLevel,
	"error":   zapcore.ErrorLevel,
	"fatal":   zapcore.FatalLevel,
}

func NewZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}

	logger.Init()
	return logger
}

func (logger *zapLogger) getLogLevel() zapcore.Level {
	level, exists := logLevelMap[logger.cfg.Logger.Level]
	if !exists {
		return zapcore.DebugLevel
	}
	return level
}

func (logger *zapLogger) Init() {

	once.Do(func() {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logger.cfg.Logger.FilePath,
			MaxSize:    1,
			MaxAge:     5,
			LocalTime:  true,
			MaxBackups: 10,
			Compress:   true,
		})
	
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
	
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			w,
			logger.getLogLevel(),
		)
	
		loggerInstance := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
		
		zapSinLogger = loggerInstance.With("AppName", "MyApp", "LoggerName", "ZapLogger")
	})
	
	logger.logger = zapSinLogger
}

func (logger *zapLogger) Info(category Category, sub SubCategory, message string, extra map[Extra]interface{}) {
	params := prepareLogKeys(extra, category, sub)

	logger.logger.Infow(message, params...)
}

func (logger *zapLogger) Debug(category Category, sub SubCategory, message string, extra map[Extra]interface{}) {
	params := prepareLogKeys(extra, category, sub)

	logger.logger.Debugw(message, params...)
}

func (logger *zapLogger) Warn(category Category, sub SubCategory, message string, extra map[Extra]interface{}) {
	params := prepareLogKeys(extra, category, sub)

	logger.logger.Warnw(message, params...)
}

func (logger *zapLogger) Error(category Category, sub SubCategory, message string, extra map[Extra]interface{}) {
	params := prepareLogKeys(extra, category, sub)

	logger.logger.Errorw(message, params...)
}

func (logger *zapLogger) Fatal(category Category, sub SubCategory, message string, extra map[Extra]interface{}) {
	params := prepareLogKeys(extra, category, sub)

	logger.logger.Fatalw(message, params...)
}

func prepareLogKeys(extra map[Extra]interface{}, category Category, sub SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[Extra]interface{}, 0)
	}
	extra["category"] = category
	extra["sub"] = sub
	params := mapToZapParams(extra)
	return params
}
