package logs

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogMode string

const (
	DevMode     LogMode = "dev"
	ProduceMode LogMode = "produce"
)

func InitLogger(fileName string, level zapcore.LevelEnabler, mode ...LogMode) *zap.SugaredLogger {
	var logMode LogMode
	if len(mode) == 0 {
		logMode = ProduceMode
	} else {
		logMode = mode[0]
	}
	var logger *zap.SugaredLogger
	switch logMode {
	case DevMode:
		dev, err := zap.NewDevelopment()
		if err != nil {
			panic("failed to init development logger")
		}
		logger = dev.Sugar()
	case ProduceMode:
		writeSyncer := getLogWriter(fileName)
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writeSyncer, level)
		logr := zap.New(core, zap.AddCaller())
		logger = logr.Sugar()
	}
	return logger

}

func getEncoder() zapcore.Encoder {
	logConf := zap.NewProductionEncoderConfig()
	logConf.EncodeTime = zapcore.ISO8601TimeEncoder
	logConf.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(logConf)
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("failed to create log file", fileName)
	}
	return zapcore.AddSync(file)
}

func GetLevelEnabler() zap.AtomicLevel {
	return zap.NewAtomicLevel()
}
