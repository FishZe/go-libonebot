package protocol

// Event 事件
//
// Reference: https://12.onebot.dev/connect/data-protocol/event/
type Event struct {
	// ID 事件ID
	ID string `json:"id"`
	// Time 事件发生的时间戳 Unix时间戳 单位为秒
	Time float64 `json:"time"`
	// Type 事件类型 meta/message/notice/request
	// 对应ActionTypeXXX
	Type string `json:"type"`
	// DetailType 事件详细类型
	DetailType string `json:"detail_type"`
	// SubType 事件子类型 详细类型的子类型
	SubType string `json:"sub_type"`
	// Self 事件发起Bot
	Self Self `json:"self"`
}

// MetaEventConnect 连接事件
//
// Reference: https://12.onebot.dev/interface/meta/events/#metaconnect
type MetaEventConnect struct {
	*Event
	// Version 机器人版本信息
	//
	// resp[get_version] OneBot实现端版本信息, 与get_version动作响应数据一致
	Version struct {
		// Impl OneBot实现端名称
		Impl string `json:"impl"`
		// Version OneBot实现端版本
		Version string `json:"version"`
		// OnebotVersion OneBot协议版本
		OnebotVersion string `json:"onebot_version"`
	} `json:"version"`
}

// NewMetaEventConnect 创建一个连接事件
func NewMetaEventConnect() (e *MetaEventConnect) {
	e = new(MetaEventConnect)
	e.Event = new(Event)
	e.Event.Type = ActionTypeMeta
	e.Event.DetailType = "connect"
	return
}

// MetaEventHeartbeat 心跳事件
type MetaEventHeartbeat struct {
	*Event
	// Interval 心跳间隔
	Interval int `json:"interval"`
}

// NewMetaEventHeartbeat 创建一个心跳事件
func NewMetaEventHeartbeat() (e *MetaEventHeartbeat) {
	e = new(MetaEventHeartbeat)
	e.Event = new(Event)
	e.Event.Type = ActionTypeMeta
	e.Event.DetailType = "heartbeat"
	return
}

// MetaEventStatusUpdate 机器人状态更新事件
//
// Reference: https://12.onebot.dev/interface/meta/events/#metastatus_update
type MetaEventStatusUpdate struct {
	*Event
	// Status 机器人状态
	Status struct {
		// Good
		Good bool `json:"good"`
		// Bots
		Bots []struct {
			// Self
			Self Self `json:"self"`
			// Online
			Online bool `json:"online"`
		} `json:"bots"`
	} `json:"status"`
}

// NewMetaEventStatusUpdate 创建一个机器人状态更新事件
func NewMetaEventStatusUpdate() (e *MetaEventStatusUpdate) {
	e = new(MetaEventStatusUpdate)
	e.Event = new(Event)
	e.Event.Type = ActionTypeMeta
	e.Event.DetailType = "status_update"
	return
}

// MessageEventPrivate 私聊消息事件
//
// Reference: https://12.onebot.dev/interface/user/message-events/#messageprivate
type MessageEventPrivate struct {
	*Event
	// MessageID 消息ID
	MessageId string `json:"message_id"`
	// Message 消息内容 消息段列表
	Message []Segment `json:"message"`
	// AltMessage 消息内容的替代表示, 可以为空
	AltMessage string `json:"alt_message"`
	// UserId 用户 ID
	UserId string `json:"user_id"`
}

// NewMessageEventPrivate 创建一个私聊消息事件
func NewMessageEventPrivate() (e *MessageEventPrivate) {
	e = new(MessageEventPrivate)
	e.Event = new(Event)
	e.Event.Type = ActionTypeMessage
	e.Event.DetailType = "private"
	return
}

// MessageEventGroup 群消息事件
//
// Reference: https://12.onebot.dev/interface/group/message-events/#messagegroup
type MessageEventGroup struct {
	*Event
	// MessageID 消息ID
	MessageId string `json:"message_id"`
	// Message 消息内容 消息段列表
	Message []Segment `json:"message"`
	// AltMessage 消息内容的替代表示, 可以为空
	AltMessage string `json:"alt_message"`
	// GroupId 群组ID
	GroupId string `json:"group_id"`
	// UserId 用户 ID
	UserId string `json:"user_id"`
}

// NewMessageEventGroup 创建一个群消息事件
func NewMessageEventGroup() (e *MessageEventGroup) {
	e = new(MessageEventGroup)
	e.Event = new(Event)
	e.Event.Type = ActionTypeMessage
	e.Event.DetailType = "group"
	return
}

// MessageEventChannel 频道消息事件
//
// Reference: https://12.onebot.dev/interface/guild/message-events/#messagechannel
type MessageEventChannel struct {
	*Event
	// MessageID 消息ID
	MessageId string `json:"message_id"`
	// Message 消息内容 消息段列表
	Message []Segment `json:"message"`
	// AltMessage 消息内容的替代表示, 可以为空
	AltMessage string `json:"alt_message"`
	// GuildId   群组ID
	GuildId string `json:"guild_id"`
	// ChannelId 频道ID
	ChannelId string `json:"channel_id"`
	// UserId    用户ID
	UserId string `json:"user_id"`
}

// NewMessageEventChannel 创建一个频道消息事件
func NewMessageEventChannel() (e *MessageEventChannel) {
	e = new(MessageEventChannel)
	e.Event = new(Event)
	e.Event.Type = ActionTypeMessage
	e.Event.DetailType = "channel"
	return
}

// NoticeEventFriendIncrease 好友增加事件
//
// Reference: https://12.onebot.dev/interface/user/notice-events/#noticefriend_increase
type NoticeEventFriendIncrease struct {
	*Event
	// UserId 用户ID
	UserId string `json:"user_id"`
}

// NewNoticeEventFriendIncrease 创建一个好友增加事件
func NewNoticeEventFriendIncrease() (e *NoticeEventFriendIncrease) {
	e = new(NoticeEventFriendIncrease)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "friend_increase"
	return
}

// NoticeEventFriendDecrease 好友减少事件
//
// Reference: https://12.onebot.dev/interface/user/notice-events/#noticefriend_decrease
type NoticeEventFriendDecrease struct {
	*Event
	// UserId 用户ID
	UserId string `json:"user_id"`
}

// NewNoticeEventFriendDecrease 创建一个好友减少事件
func NewNoticeEventFriendDecrease() (e *NoticeEventFriendDecrease) {
	e = new(NoticeEventFriendDecrease)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "friend_decrease"
	return
}

// NoticeEventMessageDelete 消息删除事件
//
// Reference: https://12.onebot.dev/interface/user/notice-events/#noticeprivate_message_delete
type NoticeEventMessageDelete struct {
	*Event
	// MessageID 消息ID
	MessageId string `json:"message_id"`
	// UserId 用户ID
	UserId string `json:"user_id"`
}

// NewNoticeEventMessageDelete 创建一个消息删除事件
func NewNoticeEventMessageDelete() (e *NoticeEventMessageDelete) {
	e = new(NoticeEventMessageDelete)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "private_message_delete"
	return
}

// NoticeEventGroupMemberIncrease 群成员增加事件
//
// Reference: https://12.onebot.dev/interface/group/notice-events/#noticegroup_member_increase
type NoticeEventGroupMemberIncrease struct {
	*Event
	// UserId 用户ID
	UserId string `json:"user_id"`
	// GroupId 群组ID
	GroupId string `json:"group_id"`
	// OperatorId 操作者ID
	OperatorId string `json:"operator_id"`
}

// NewNoticeEventGroupMemberIncrease 创建一个群成员增加事件
//
// subType: join / invite / 自定义 / 空
func NewNoticeEventGroupMemberIncrease(subType string) (e *NoticeEventGroupMemberIncrease) {
	e = new(NoticeEventGroupMemberIncrease)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "group_member_increase"
	e.Event.SubType = subType
	return
}

// NoticeEventGroupMemberDecrease 群成员减少事件
//
// Reference: https://12.onebot.dev/interface/group/notice-events/#noticegroup_member_decrease
type NoticeEventGroupMemberDecrease struct {
	*Event
	// UserId 用户ID
	UserId string `json:"user_id"`
	// GroupId 群组ID
	GroupId string `json:"group_id"`
	// OperatorId 操作者ID
	OperatorId string `json:"operator_id"`
}

// NewNoticeEventGroupMemberDecrease 创建一个群成员减少事件
//
// subType: leave / kick / 自定义 / 空
func NewNoticeEventGroupMemberDecrease(subType string) (e *NoticeEventGroupMemberDecrease) {
	e = new(NoticeEventGroupMemberDecrease)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "group_member_decrease"
	e.Event.SubType = subType
	return
}

// NoticeEventGroupMessageDelete 消息删除事件
//
// Reference: https://12.onebot.dev/interface/group/notice-events/#noticegroup_message_delete
type NoticeEventGroupMessageDelete struct {
	*Event
	// MessageId 消息ID
	MessageId string `json:"message_id"`
	// UserId 用户ID
	UserId string `json:"user_id"`
	// OperatorId 群组ID
	OperatorId string `json:"operator_id"`
}

// NewNoticeEventGroupMessageDelete 创建一个群消息删除事件
//
// subType: recall / delete / 自定义 / 空
func NewNoticeEventGroupMessageDelete(subType string) (e *NoticeEventGroupMessageDelete) {
	e = new(NoticeEventGroupMessageDelete)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "group_message_delete"
	e.Event.SubType = subType
	return
}

// NoticeEventGuildMemberIncrease 群组成员增加事件
//
// Reference: https://12.onebot.dev/interface/guild/notice-events/#noticeguild_member_increase
type NoticeEventGuildMemberIncrease struct {
	*Event
	// UserId 用户ID
	UserId string `json:"user_id"`
	// GuildId 群组ID
	GuildId string `json:"guild_id"`
	// OperatorId 操作者ID
	OperatorId string `json:"operator_id"`
}

// NewNoticeEventGuildMemberIncrease 创建一个群组成员增加事件
//
// subType: join / invite / 自定义 / 空
func NewNoticeEventGuildMemberIncrease(subType string) (e *NoticeEventGuildMemberIncrease) {
	e = new(NoticeEventGuildMemberIncrease)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "guild_member_increase"
	e.Event.SubType = subType
	return
}

// NoticeEventGuildMemberDecrease 群组成员减少事件
//
// Reference: https://12.onebot.dev/interface/guild/notice-events/#noticeguild_member_decrease
type NoticeEventGuildMemberDecrease struct {
	*Event
	// UserId 用户ID
	UserId string `json:"user_id"`
	// GuildId 群组ID
	GuildId string `json:"guild_id"`
	// OperatorId 操作者ID
	OperatorId string `json:"operator_id"`
}

// NewNoticeEventGuildMemberDecrease 创建一个群组成员减少事件
//
// subType: leave / kick / 自定义 / 空
func NewNoticeEventGuildMemberDecrease(subType string) (e *NoticeEventGuildMemberDecrease) {
	e = new(NoticeEventGuildMemberDecrease)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "guild_member_decrease"
	e.Event.SubType = subType
	return
}

// NoticeEventChannelMemberIncrease 频道成员增加事件
//
// Reference: https://12.onebot.dev/interface/guild/notice-events/#noticechannel_member_increase
type NoticeEventChannelMemberIncrease struct {
	*Event
	// UserId 用户ID
	UserId string `json:"user_id"`
	// GuildId 群组ID
	GuildId string `json:"guild_id"`
	// ChannelId  频道ID
	ChannelId string `json:"channel_id"`
	// OperatorId 操作者ID
	OperatorId string `json:"operator_id"`
}

// NewNoticeEventChannelMemberIncrease 创建一个频道成员增加事件
//
// subType: join / invite / 自定义 / 空
func NewNoticeEventChannelMemberIncrease(subType string) (e *NoticeEventChannelMemberIncrease) {
	e = new(NoticeEventChannelMemberIncrease)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "channel_member_increase"
	e.Event.SubType = subType
	return
}

// NoticeEventChannelMemberDecrease 频道成员减少事件
//
// Reference: https://12.onebot.dev/interface/guild/notice-events/#noticechannel_member_decrease
type NoticeEventChannelMemberDecrease struct {
	*Event
	// UserId 用户ID
	UserId string `json:"user_id"`
	// GuildId 群组ID
	GuildId string `json:"guild_id"`
	// ChannelId  频道ID
	ChannelId string `json:"channel_id"`
	// OperatorId 操作者ID
	OperatorId string `json:"operator_id"`
}

// NewNoticeEventChannelMemberDecrease 创建一个频道成员减少事件
//
// subType: leave / kick / 自定义 / 空
func NewNoticeEventChannelMemberDecrease(subType string) (e *NoticeEventChannelMemberDecrease) {
	e = new(NoticeEventChannelMemberDecrease)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "channel_member_decrease"
	e.Event.SubType = subType
	return
}

// NoticeEventChannelMessageDelete 消息删除事件
//
// Reference: https://12.onebot.dev/interface/guild/notice-events/#noticechannel_message_delete
type NoticeEventChannelMessageDelete struct {
	*Event
	// GuildId 群组ID
	GuildId string `json:"guild_id"`
	// ChannelId  频道ID
	ChannelId string `json:"channel_id"`
	// MessageId 消息ID
	MessageId string `json:"message_id"`
	// UserId 用户ID
	UserId string `json:"user_id"`
	// OperatorId 操作者ID
	OperatorId string `json:"operator_id"`
}

// NewNoticeEventChannelMessageDelete 创建一个消息删除事件
//
// subType: recall / delete / 自定义 / 空
func NewNoticeEventChannelMessageDelete(subType string) (e *NoticeEventChannelMessageDelete) {
	e = new(NoticeEventChannelMessageDelete)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "channel_message_delete"
	e.Event.SubType = subType
	return
}

// NoticeEventChannelCreate 频道创建事件
//
// Reference: https://12.onebot.dev/interface/guild/notice-events/#noticechannel_create
type NoticeEventChannelCreate struct {
	*Event
	// GuildId 群组ID
	GuildId string `json:"guild_id"`
	// ChannelId 频道ID
	ChannelId string `json:"channel_id"`
	// OperatorId 操作者ID
	OperatorId string `json:"operator_id"`
}

// NewNoticeEventChannelCreate 创建一个频道创建事件
func NewNoticeEventChannelCreate() (e *NoticeEventChannelCreate) {
	e = new(NoticeEventChannelCreate)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "channel_create"
	return
}

// NoticeEventChannelDelete 频道删除事件
//
// Reference: https://12.onebot.dev/interface/guild/notice-events/#noticechannel_delete
type NoticeEventChannelDelete struct {
	*Event
	// GuildId 群组ID
	GuildId string `json:"guild_id"`
	// ChannelId 频道ID
	ChannelId string `json:"channel_id"`
	// OperatorId 操作者ID
	OperatorId string `json:"operator_id"`
}

// NewNoticeEventChannelDelete 创建一个频道删除事件
func NewNoticeEventChannelDelete() (e *NoticeEventChannelDelete) {
	e = new(NoticeEventChannelDelete)
	e.Event = new(Event)
	e.Event.Type = ActionTypeNotice
	e.Event.DetailType = "channel_delete"
	return
}
