package util

import (
	"fmt"
)

// LogLevel 日志等级
var (
	LogLevelInfo    = 1 << 0
	LogLevelError   = LogLevelInfo | (1 << 1)
	LogLevelWarning = LogLevelError | (1 << 2)
	LogLevelDebug   = LogLevelWarning | (1 << 3)
)

// 各种颜色
var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// LogInterface 日志接口
type LogInterface interface {
	Println(string)
	Info(string)
	Warning(string)
	Error(string)
	Debug(string)
	SetLogLevel(int)
}

// Logger log的接口
var Logger LogInterface

// SetLogger 设置日志Logger
func SetLogger(l LogInterface) {
	Logger = l
}

// Log Logger的默认实现
type Log struct {
	level int
}

// SetLogLevel 修改日志等级
func (l *Log) SetLogLevel(level int) {
	l.level = level
}

// getColor 获取颜色
func (l *Log) getColor(level int) string {
	switch level {
	case LogLevelInfo:
		return Green
	case LogLevelWarning:
		return Yellow
	case LogLevelError:
		return Red
	case LogLevelDebug:
		return Blue
	default:
		return Reset
	}
}

// getLevelString 获取等级字符串
func (l *Log) getLevelString(level int) string {
	switch level {
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarning:
		return "WARNING"
	case LogLevelError:
		return "ERROR"
	case LogLevelDebug:
		return "DEBUG"
	default:
		return "PRINT"
	}
}

// print 打印
func (l *Log) print(level int, s string) {
	if l.level >= level {
		fmt.Println(fmt.Sprintf("%s[LIBOB][%s][%s]%s%s", l.getColor(level), GetFormatTime(), l.getLevelString(level), s, l.getColor(-1)))
	}
}

// Info 打印Info
func (l *Log) Info(s string) {
	l.print(LogLevelInfo, s)
}

// Warning 打印Warning
func (l *Log) Warning(s string) {
	l.print(LogLevelWarning, s)
}

// Error 打印Error
func (l *Log) Error(s string) {
	l.print(LogLevelError, s)
}

// Debug 打印Debug
func (l *Log) Debug(s string) {
	l.print(LogLevelDebug, s)
}

// Println 打印
func (l *Log) Println(s string) {
	l.print(0, s)
}

func init() {
	// 使用默认日志
	SetLogger(new(Log))
	// 等级
	Logger.SetLogLevel(LogLevelWarning)
}
