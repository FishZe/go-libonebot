package go_libonebot

// OneBotConfig OneBot协议配置
type OneBotConfig struct {
	// PlatForm 平台
	PlatForm string `json:"platform"`
	// Implementation 实现
	Implementation string `json:"implementation"`
	// Version 版本
	Version string `json:"version"`
}
