package v11_adapter

// MetaEventLifeCycle  元事件生命周期
type MetaEventLifeCycle struct {
	*V11EventHead
}

// MetaEventHeartBeat 元事件心跳
type MetaEventHeartBeat struct {
	*V11EventHead
	// Status 机器人状态
	Status struct {
		// Online 机器人是否在线
		Online bool `json:"online"`
		// Good 机器人是否健康
		Good bool `json:"good"`
	}
	// Interval 心跳间隔
	Interval int `json:"interval"`
}
