package Middleware

import (
	"bytes"
	Config "gin-laravel/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

var logger *zap.SugaredLogger

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	//logFilePath := Config.App.GetString("")
	//logFileName := Config.App.GetString("")
	logFilePath := Config.App.GetString("server.LogPath") + "/"
	logFileName := time.Now().Format("2006-01-02") + ".log"
	loglevel := "info"

	hook := lumberjack.Logger{
		Filename:   logFilePath + logFileName, // ⽇志⽂件路径
		MaxSize:    1024,                      // megabytes
		MaxBackups: 3,                         // 最多保留3个备份
		MaxAge:     365,                       //days
		Compress:   true,                      // 是否压缩 disabled by default
	}
	fileWriter := zapcore.AddSync(&hook)
	var highPriority zapcore.Level
	switch loglevel {
	case "debug":
		highPriority = zap.DebugLevel
	case "info":
		highPriority = zap.InfoLevel
	case "error":
		highPriority = zap.ErrorLevel
	default:
		highPriority = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//core := zapcore.NewCore(
	//	zapcore.NewConsoleEncoder(encoderConfig),
	//	fileWriter,
	//	level,
	//)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleDebugging := zapcore.Lock(os.Stdout)
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})
	core := zapcore.NewTee(
		// 打印在控制台
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		// 打印在文件中
		zapcore.NewCore(consoleEncoder, fileWriter, highPriority),
	)

	//代码的位置也可以输出
	//logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = zap.New(core).Sugar()

	return func(c *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		body := ""

		if w.body.Len() < 1000 {
			body = w.body.String()
		}

		logger.Infow(c.Request.Method,
			"uri", c.Request.RequestURI,
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"time", latencyTime,
			"ip", c.ClientIP(),
			"body", body,
			"size", w.body.Len(),
		)
	}
}
