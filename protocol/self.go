package protocol

// Self 字段 用于区分不同机器人
type Self struct {
	// PlatForm 平台
	PlatForm string `json:"platform"`
	// UserId 用户ID
	UserId string `json:"user_id"`
}
