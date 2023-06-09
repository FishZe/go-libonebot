package util

import (
	"github.com/google/uuid"
	"net"
	"strconv"
	"time"
)

// GetUUID 获取一个随机的UUID
//
// 来自github.com/google/uuid
func GetUUID() string {
	id := uuid.New()
	return id.String()
}

// GetTimeStampFloat64 获取当前时间戳
//
// https://12.onebot.dev/connect/data-protocol/event/ 要求的Time为float64
func GetTimeStampFloat64() float64 {
	return float64(time.Now().UnixMilli()) / 1000
}

// GetTimeStamp 获取当前时间戳 秒
// 备用
func GetTimeStamp() int64 {
	return time.Now().Unix()
}

// GetNanoTimeStamp 获取当前时间戳 纳秒
func GetNanoTimeStamp() int64 {
	return time.Now().UnixNano()
}

// GetFormatTime 获取当前时间的格式化字符串
func GetFormatTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// CheckPortAvailable 检查端口是否可用
func CheckPortAvailable(port int) bool {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return false
	}
	err = ln.Close()
	if err != nil {
		return false
	}
	return true
}
