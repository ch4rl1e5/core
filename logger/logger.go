package logger

import (
	"github.com/ch4rl1e5/go-common/constants"
	"go.uber.org/zap"
)

var ZapLogger *zap.Logger

func StartupLogger(config Config) {
	var err error
	if config.Enviroment == constants.ProdEnvironment {
		ZapLogger, err = zap.NewProduction()
	}

	if config.Enviroment == constants.DevEnvironment {
		ZapLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}
}
