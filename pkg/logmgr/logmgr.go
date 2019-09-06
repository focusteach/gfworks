package logmgr

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logPath string

var (
	singleton *LoggerMgr
	once      sync.Once
)

//GetInstance 用于获取单例模式对象
func GetInstance() *LoggerMgr {
	once.Do(func() {
		singleton = &LoggerMgr{}
		singleton.init()
	})

	return singleton
}

// LOGLEVEL 日志级别
type LOGLEVEL int

const (
	_ LOGLEVEL = iota
	TraceLevel
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

func init() {

	flag.StringVar(&logPath, "logs", "./logs/", "default config path")

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := os.MkdirAll(logPath, 0777); err != nil {
		log.Fatalf("conf path make failed.err：%+v", err)
	}
}

// Logger logger
type Logger struct {
	logger *logrus.Logger
	Name   string
}

// LoggerMgr log管理器
type LoggerMgr struct {
	logs map[string]*Logger
}

func (mgr *LoggerMgr) init() {
	mgr.logs = make(map[string]*Logger)
}

// GetLogger 获得一个logger
func (mgr *LoggerMgr) GetLogger(logName string) *Logger {
	// 默认文件名
	if logName == "" {
		logName = "app"
	}
	if logger, ok := mgr.logs[logName]; ok {
		return logger
	}

	return mgr.InitLogger(logName)
}

// InitLogger 初始化访问日志
func (mgr *LoggerMgr) InitLogger(logName string) *Logger {

	var logInstance Logger
	logInstance.Name = logName

	logger := logrus.New()
	fileName := path.Join(logPath, fmt.Sprintf("%s.log", logName))
	logPath, _ := filepath.Abs(fileName)

	//禁止logrus的输出
	src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|0777)
	if err != nil {
		fmt.Println("err", err)
	}
	// 设置日志输出的路径
	logger.Out = src
	logger.SetLevel(logrus.TraceLevel)
	//apiLogPath := "gin-api.log"
	logWriter, err := rotatelogs.New(
		logPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(logPath),          // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logger.AddHook(lfHook)

	logInstance.logger = logger
	mgr.logs[logName] = &logInstance
	return &logInstance
}

// Dir log dir
func Dir() string {
	return logPath
}

// Logf 格式化输出
func Logf(logName string, logLevel LOGLEVEL, format string, args ...interface{}) {
	logClient := GetInstance().GetLogger(logName)

	switch logLevel {
	case TraceLevel:
		logClient.logger.Tracef(format, args...)
	case DebugLevel:
		logClient.logger.Debugf(format, args...)
	case InfoLevel:
		logClient.logger.Infof(format, args...)
	case WarnLevel:
		logClient.logger.Warnf(format, args...)
	case ErrorLevel:
		logClient.logger.Errorf(format, args...)
	case PanicLevel:
		logClient.logger.Panicf(format, args...)
	case FatalLevel:
		logClient.logger.Fatalf(format, args...)
	}
}

// Logln 换行输出
func Logln(logName string, logLevel LOGLEVEL, args ...interface{}) {
	logClient := GetInstance().GetLogger(logName)

	switch logLevel {
	case TraceLevel:
		logClient.logger.Traceln(args...)
	case DebugLevel:
		logClient.logger.Debugln(args...)
	case InfoLevel:
		logClient.logger.Infoln(args...)
	case WarnLevel:
		logClient.logger.Warnln(args...)
	case ErrorLevel:
		logClient.logger.Errorln(args...)
	case PanicLevel:
		logClient.logger.Panicln(args...)
	case FatalLevel:
		logClient.logger.Fatalln(args...)
	}
}
