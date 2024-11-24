import (
    "gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
    logger := &lumberjack.Logger{
        Filename:   "./logs/app.log",
        MaxSize:    5,  // Megabytes
        MaxBackups: 3,
        MaxAge:     28, // Days
    }
    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
        zapcore.AddSync(logger),
        zap.InfoLevel,
    )
    Logger = zap.New(core)
}
