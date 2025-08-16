package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	log  *zap.Logger
	once sync.Once
)

// Initialize the global logger
func Init(env string) {
	once.Do(func() {
		var err error
		if env == "dev" {
			log, err = zap.NewDevelopment()
		} else {
			log, err = zap.NewProduction()
		}

		if err != nil {
			panic(err)
		}
	})
}

func L() *zap.Logger {
	if log == nil {
		panic("logger not initialized - Init() must be called before")
	}
	return log
}
