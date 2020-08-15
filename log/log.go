package log

/**
log日志相关
*/

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log = logrus.New() // 创建一个log实例

/**
log文件生成路径
当天日志文件路径
*/
var LogFilePath = "./cyan.log"

/**
历史log文件保存路径
*/
var LogFileSavePath = "./logging/cyan"

/**
历史log文件保存格式后缀
示例: cyan.20200519.log
*/
var LogFileSuffix = ".%Y%m%d.log"

/**
设置最大保存时间(默认180天)
*/
var WithMaxAge = 180 * 24 * time.Hour

/**
设置日志切割时间间隔(默认1天)
*/
var WithRotationTime = 24 * time.Hour

/**
初始化日志文件
搭配gin框架使用
*/
func IntiLog() {
	Log.Formatter = &logrus.JSONFormatter{}                                       // 设置为json格式的日志
	f, err := os.OpenFile(LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) // 创建一个log日志文件
	if err != nil {
		panic(err)
	}

	Log.Out = f                   // 设置log的默认文件输出
	gin.SetMode(gin.ReleaseMode)  // 线上模式，控制台不会打印信息
	gin.DefaultWriter = Log.Out   // gin框架自己记录的日志也会输出
	Log.Level = logrus.DebugLevel // 设置日志级别

	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		LogFileSavePath+LogFileSuffix,
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(LogFileSavePath),
		// 设置最大保存时间(180天)
		rotatelogs.WithMaxAge(WithMaxAge),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(WithRotationTime),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	Log.AddHook(lfHook)

	Log.Info("Cyan start.....")
}
