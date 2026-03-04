package intermediate

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	logger, err := config.Build()
	if err != nil {
		log.Println("Error in initializing Zap logger:", err)
		return
	}

	defer logger.Sync()

	logger.Info("This is an info message.")

	logger.Info("User logged in.", zap.String("username", "John Doe"), zap.String("method", "GET"))
}
