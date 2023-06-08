package util

import (
	"github.com/google/uuid"
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

// GetTimeStamp 获取当前时间戳
// 备用
func GetTimeStamp() int64 {
	return time.Now().Unix()
}

// GetFormatTime 获取当前时间的格式化字符串
func GetFormatTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
