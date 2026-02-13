package log

import (
	"sync"

	"go.uber.org/zap"
)

var (
	base *zap.Logger
	once sync.Once
)

func InitLogger(dev bool) {
	once.Do(func() {
		if dev {
			base, _ = zap.NewDevelopment()
		} else {
			base, _ = zap.NewProduction()
		}
	})
}

func mustBase() *zap.Logger {
	if base == nil {
		panic("log.InitLogger must be called before using the logger")
	}
	return base
}

func Core(name string) *zap.SugaredLogger {
	return mustBase().
		Named("core").
		With(zap.String("module", name)).
		Sugar()
}

func Service(name string) *zap.SugaredLogger {
	return mustBase().
		Named("service").
		With(zap.String("service", name)).
		Sugar()
}

func Sync() {
	if base != nil {
		_ = base.Sync()
	}
}
