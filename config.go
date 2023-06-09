package go_libonebot

// OneBotConfig OneBot协议配置
type OneBotConfig struct {
	// PlatForm 平台
	PlatForm string `json:"platform"`
	// Implementation 实现
	Implementation string `json:"implementation"`
	// Version 版本
	Version string `json:"version"`
	// HeartBeat 心跳
	HeartBeat struct {
		// Enable 是否启用
		Enable bool `json:"enable"`
		// Interval 心跳间隔
		Interval int `json:"interval"`
	}
}
