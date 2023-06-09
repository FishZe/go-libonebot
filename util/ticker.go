package util

import "fmt"

// Ticker 计时器
type Ticker struct {
	// startTime 开始时间
	startTime int64
	// endTime 结束时间
	endTime int64
}

// NewTicker 创建一个计时器
func NewTicker() *Ticker {
	return &Ticker{
		startTime: GetNanoTimeStamp(),
	}
}

func (t *Ticker) getUint(x int) string {
	switch x {
	case 1:
		return "ns"
	case 2:
		return "us"
	case 3:
		return "ms"
	case 4:
		return "s"
	case 5:
		return "ks"
	}
	return ""
}

// Stop 停止计时
func (t *Ticker) Stop() {
	t.endTime = GetNanoTimeStamp()
}

// GetDuration 获取计时器的时间 纳秒
func (t *Ticker) GetDuration() int64 {
	return t.endTime - t.startTime
}

// GetDurationString 获取计时器的时间 字符串 包含单位
func (t *Ticker) GetDurationString() string {
	var n = t.GetDuration()
	var s = 0
	for n > 1000 && s <= 5 {
		s++
		n /= 1000
	}
	return fmt.Sprintf("%d %s", n, t.getUint(s))
}
