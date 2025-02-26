package logger

import (
	"github.com/devmyong/todo/pkg/config"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewZapLogger(cfg *config.LoggerConfig) (*zap.Logger, error) {
	// 1) Set log level
	logLevel := zap.NewAtomicLevel()
	if err := logLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, err
	}

	// 2) Set output path
	writeSyncer := zapcore.AddSync(os.Stdout)

	// 3) Set encoder
	var encoderCfg zapcore.EncoderConfig
	if cfg.Environment == config.EnvProduction {
		encoderCfg = ecszap.NewDefaultEncoderConfig().ToZapCoreEncoderConfig()
		encoderCfg.TimeKey = "ts"
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeDuration = zapcore.StringDurationEncoder
	}

	var encoder zapcore.Encoder
	if cfg.Format == config.LogFormatJSON {
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	// 4) Create core
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	// 5) Create logger
	logger := zap.New(core,
		zap.Hooks(Hook()),
		zap.AddCaller(),
		zap.Fields(
			zap.String("env", cfg.Environment),
			zap.String("version", cfg.Version),
		),
	)

	return logger, nil
}

// todo change name and implement
func Hook() func(zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		if entry.Level == zapcore.ErrorLevel {
			// todo send to slack
		}
		return nil
	}
}
