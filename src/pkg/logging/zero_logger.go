package logging

import (
	"os"
	"sync"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once
var zeroSinLogger *zerolog.Logger

type zeroLogger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

var zeroLogLevelMap = map[string]zerolog.Level{
	"debug":   zerolog.DebugLevel,
	"info":    zerolog.InfoLevel,
	"warning": zerolog.WarnLevel,
	"error":   zerolog.ErrorLevel,
	"fatal":   zerolog.FatalLevel,
}

func NewZeroLogger(cfg *config.Config) *zeroLogger {
	logger := &zeroLogger{cfg: cfg}

	logger.Init()
	return logger
}

func (logger *zeroLogger) getLogLevel() zerolog.Level {
	level, exists := zeroLogLevelMap[logger.cfg.Logger.Level]
	if !exists {
		return zerolog.DebugLevel
	}
	return level
}

func (logger *zeroLogger) Init() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		file, err := os.OpenFile(logger.cfg.Logger.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic("Couldn't open log file")
		}
		var loggerInstance = zerolog.New(file).With().Timestamp().Str("AppName", "MyApp").Str("LoggerName", "Zerolog").Logger()
		zerolog.SetGlobalLevel(logger.getLogLevel())
		zeroSinLogger = &loggerInstance 
	})
	

	logger.logger = zeroSinLogger
}

func (logger *zeroLogger) Info(category Category, sub SubCategory, message string, extra map[Extra]interface{}) {
	logger.logger.Info().Str("Category", string(category)).Str("SubCategory", string(sub)).Fields(mapToZeroParams(extra)).Msg(message)
}

func (logger *zeroLogger) Debug(category Category, sub SubCategory, message string, extra map[Extra]interface{}) {

	logger.logger.Debug().Str("Category", string(category)).Str("SubCategory", string(sub)).Fields(mapToZeroParams(extra)).Msg(message)
}

func (logger *zeroLogger) Warn(category Category, sub SubCategory, message string, extra map[Extra]interface{}) {
	logger.logger.Warn().Str("Category", string(category)).Str("SubCategory", string(sub)).Fields(mapToZeroParams(extra)).Msg(message)

}

func (logger *zeroLogger) Error(category Category, sub SubCategory, message string, extra map[Extra]interface{}) {
	logger.logger.Error().Str("Category", string(category)).Str("SubCategory", string(sub)).Fields(mapToZeroParams(extra)).Msg(message)

}

func (logger *zeroLogger) Fatal(category Category, sub SubCategory, message string, extra map[Extra]interface{}) {
	logger.logger.Fatal().Str("Category", string(category)).Str("SubCategory", string(sub)).Fields(mapToZeroParams(extra)).Msg(message)
}
