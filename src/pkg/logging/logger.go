package logging

import "github.com/erfanwd/golang-course-project/config"

type Logger interface {
	Init()
	Info(category Category, sub SubCategory, message string, extra map[Extra]interface{})
	Debug(category Category, sub SubCategory, message string, extra map[Extra]interface{})
	Warn(category Category, sub SubCategory, message string, extra map[Extra]interface{})
	Error(category Category, sub SubCategory, message string, extra map[Extra]interface{})
	Fatal(category Category, sub SubCategory, message string, extra map[Extra]interface{})
}

func NewLogger(cfg *config.Config) Logger {
	if cfg.Logger.Logger == "zap" {
		return NewZapLogger(cfg);	
	}else if cfg.Logger.Logger == "zerolog" {
		return NewZeroLogger(cfg);
	}
	panic("logger is not valid")
	
}
