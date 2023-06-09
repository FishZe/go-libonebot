package protocol

// RequestGetLatestEvents 获取最新事件列表
//
// 仅 HTTP 通信方式必须支持，用于轮询获取事件。
//
// Reference: https://12.onebot.dev/interface/meta/actions/#get_latest_events
type RequestGetLatestEvents struct {
	*Request
	// Limit 获取的事件数量上限，0 表示不限制
	Limit int64 `json:"limit"`
	// Timeout 没有事件时最多等待的秒数，0 表示使用短轮询，不等待
	Timeout int64 `json:"timeout"`
	// f 用户自定义函数
	f func(events *RequestGetLatestEvents) *ResponseGetLatestEvents
}

// New 构造函数
func (*RequestGetLatestEvents) New() any {
	return &RequestGetLatestEvents{
		Request: &Request{
			Action: "get_latest_events",
		},
	}
}

// Do 执行
//
// 执行用户自定义函数
func (r *RequestGetLatestEvents) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	resp := r.f(r)
	newEvents := make([]any, 0)
	for _, v := range resp.Events {
		if _, err := EventCheck(r.Self, v); err == nil {
			newEvents = append(newEvents, v)
		}
	}
	resp.Events = newEvents
	return resp
}

// RequestGetSupportedActions 获取支持的动作列表
//
// Reference: https://12.onebot.dev/interface/meta/actions/#get_supported_actions
type RequestGetSupportedActions struct {
	*Request
	// f 用户自定义函数
	f func(events *RequestGetSupportedActions) *ResponseGetSupportedActions
}

// New 构造函数
func (*RequestGetSupportedActions) New() any {
	return &RequestGetSupportedActions{
		Request: &Request{
			Action: "get_supported_actions",
		},
	}
}

// Do 执行
//
// 执行用户自定义函数
func (r *RequestGetSupportedActions) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetStatus  获取运行状态
//
// Reference: https://12.onebot.dev/interface/meta/actions/#get_status
type RequestGetStatus struct {
	*Request
	// f 用户自定义函数
	f func(events *RequestGetStatus) *ResponseGetStatus
}

// New 构造函数
func (*RequestGetStatus) New() any {
	return &RequestGetStatus{
		Request: &Request{
			Action: "get_status",
		},
	}
}

// Do 执行
//
// 执行用户自定义函数
func (r *RequestGetStatus) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetVersion 获取版本信息
//
// Reference: https://12.onebot.dev/interface/meta/actions/#get_version
type RequestGetVersion struct {
	*Request
	// f 用户自定义函数
	f func(events *RequestGetVersion) *ResponseGetVersion
}

// New 构造函数
func (*RequestGetVersion) New() any {
	return &RequestGetVersion{
		Request: &Request{
			Action: "get_version",
		},
	}
}

// Do 执行
//
// 执行用户自定义函数
func (r *RequestGetVersion) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestSendMessage 发送消息
//
// 对于不同平台的 detail_type，如果符合标准所定义的类型，如私聊对应 private、群组对应 group，则建议使用标准定义的 detail_type 和 xxx_id。
//
// 对于其它具体类型，例如过去 QQ 还存在讨论组的情况，可以指定 detail_type 为 qq.discuss，然后参数使用 qq.discuss_id 指示讨论组 ID。
//
// 更多详细扩展规则请参考 扩展规则 https://12.onebot.dev/interface/rules/。
//
// Reference: https://12.onebot.dev/interface/message/actions/#send_message
type RequestSendMessage struct {
	*Request
	// DetailType 发送的类型，可以为 private、group、channel 或扩展的类型，和消息事件的 detail_type 字段对应
	DetailType string `json:"detail_type"`
	// UserId 用户 ID，当 detail_type 为 private 时必须传入
	UserId string `json:"user_id"`
	// GroupId	群 ID，当 detail_type 为 group 时必须传入
	GroupId string `json:"group_id"`
	// GuildId Guild 群组 ID，当 detail_type 为 channel 时必须传入
	GuildId string `json:"guild_id"`
	// ChannelId 频道 ID，当 detail_type 为 channel 时必须传入
	ChannelId string `json:"channel_id"`
	// Message 消息段列表 消息内容
	Message []Segment `json:"message"`
	// f 用户自定义函数
	f func(events *RequestSendMessage) *ResponseSendMessage
}

// New 构造函数
func (*RequestSendMessage) New() any {
	return &RequestSendMessage{
		Request: &Request{
			Action: "send_message",
		},
	}
}

// Do 执行
//
// 执行用户自定义函数
func (r *RequestSendMessage) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestDeleteMessage 撤回消息
//
// Reference: https://12.onebot.dev/interface/message/actions/#delete_message
type RequestDeleteMessage struct {
	*Request
	// MessageId 唯一的消息 ID
	MessageId string `json:"message_id"`
	// f 用户自定义函数
	f func(events *RequestDeleteMessage) *ResponseDeleteMessage
}

// New 构造函数
func (*RequestDeleteMessage) New() any {
	return &RequestDeleteMessage{
		Request: &Request{
			Action: "delete_message",
		},
	}
}

// Do 执行
//
// 执行用户自定义函数
func (r *RequestDeleteMessage) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetSelfInfo 获取机器人自身信息
//
// Reference: https://12.onebot.dev/interface/user/actions/#get_self_info
type RequestGetSelfInfo struct {
	*Request
	// f
	f func(events *RequestGetSelfInfo) *ResponseGetSelfInfo
}

// Do 执行
func (r *RequestGetSelfInfo) Do() (e any) {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// New 构造函数
func (*RequestGetSelfInfo) New() any {
	return &RequestGetSelfInfo{
		Request: &Request{
			Action: "get_self_info",
		},
	}
}

// RequestGetUserInfo 获取用户信息
//
// Reference: https://12.onebot.dev/interface/user/actions/#get_user_info
type RequestGetUserInfo struct {
	*Request
	// UserId 用户 ID，可以是好友，也可以是陌生人
	UserId string `json:"user_id"`
	// f
	f func(events *RequestGetUserInfo) *ResponseGetUserInfo
}

// Do 执行
func (r *RequestGetUserInfo) Do() (e any) {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// New 构造函数
func (*RequestGetUserInfo) New() any {
	return &RequestGetUserInfo{
		Request: &Request{
			Action: "get_user_info",
		},
	}
}

// RequestGetFriendList 获取好友列表
//
// 获取机器人的关注者或好友列表。
//
// Reference: https://12.onebot.dev/interface/user/actions/#get_friend_list
type RequestGetFriendList struct {
	*Request
	// f
	f func(events *RequestGetFriendList) *ResponseGetFriendList
}

// New 构造函数
func (*RequestGetFriendList) New() any {
	return &RequestGetFriendList{
		Request: &Request{
			Action: "get_friend_list",
		},
	}
}

// Do  执行
func (r *RequestGetFriendList) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetGroupInfo 获取群信息
//
// Reference: https://12.onebot.dev/interface/group/actions/#get_group_info
type RequestGetGroupInfo struct {
	*Request
	// GroupId 	群 ID
	GroupId string `json:"group_id"`
	// f
	f func(events *RequestGetGroupInfo) *ResponseGetGroupInfo
}

// New 构造函数
func (*RequestGetGroupInfo) New() any {
	return &RequestGetGroupInfo{
		Request: &Request{
			Action: "get_group_info",
		},
	}
}

// Do 执行
func (r *RequestGetGroupInfo) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetGroupList 获取群列表
//
// Reference: https://12.onebot.dev/interface/group/actions/#get_group_list
type RequestGetGroupList struct {
	*Request
	// f
	f func(events *RequestGetGroupList) *ResponseGetGroupList
}

// New 构造函数
func (*RequestGetGroupList) New() any {
	return &RequestGetGroupList{
		Request: &Request{
			Action: "get_group_list",
		},
	}
}

// Do 执行
func (r *RequestGetGroupList) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetGroupMemberInfo 获取群成员信息
//
// Reference: https://12.onebot.dev/interface/group/actions/#get_group_member_info
type RequestGetGroupMemberInfo struct {
	*Request
	// GroupId 	群 ID
	GroupId string `json:"group_id"`
	// UserId 	用户 ID
	UserId string `json:"user_id"`
	// f
	f func(events *RequestGetGroupMemberInfo) *ResponseGetGroupMemberInfo
}

// New 构造函数
func (*RequestGetGroupMemberInfo) New() any {
	return &RequestGetGroupMemberInfo{
		Request: &Request{
			Action: "get_group_member_info",
		},
	}
}

// Do 执行
func (r *RequestGetGroupMemberInfo) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetGroupMemberList 获取群成员列表
//
// Reference: https://12.onebot.dev/interface/group/actions/#get_group_member_list
type RequestGetGroupMemberList struct {
	*Request
	// GroupId 	群 ID
	GroupId string `json:"group_id"`
	// f
	f func(events *RequestGetGroupMemberList) *ResponseGetGroupMemberList
}

// New 构造函数
func (*RequestGetGroupMemberList) New() any {
	return &RequestGetGroupMemberList{
		Request: &Request{
			Action: "get_group_member_list",
		},
	}
}

// Do 执行
func (r *RequestGetGroupMemberList) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestSetGroupName 设置群名称
//
// Reference: https://12.onebot.dev/interface/group/actions/#set_group_name
type RequestSetGroupName struct {
	*Request
	// GroupId 群 ID
	GroupId string `json:"group_id"`
	// GroupName 新群名称
	GroupName string `json:"group_name"`
	// f
	f func(events *RequestSetGroupName) *ResponseSetGroupName
}

// New 构造函数
func (*RequestSetGroupName) New() any {
	return &RequestSetGroupName{
		Request: &Request{
			Action: "set_group_name",
		},
	}
}

// Do 执行
func (r *RequestSetGroupName) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestLeaveGroup 退出群
//
// Reference: https://12.onebot.dev/interface/group/actions/#leave_group
type RequestLeaveGroup struct {
	*Request
	// GroupId 群 ID
	GroupId string `json:"group_id"`
	// f
	f func(events *RequestLeaveGroup) *ResponseLeaveGroup
}

// New 构造函数
func (*RequestLeaveGroup) New() any {
	return &RequestLeaveGroup{
		Request: &Request{
			Action: "leave_group",
		},
	}
}

// Do 执行
func (r *RequestLeaveGroup) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetGuildInfo 获取群组信息
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_guild_info
type RequestGetGuildInfo struct {
	*Request
	// GuildIf 群组 ID
	GuildId string `json:"guild_id"`
	// f
	f func(events *RequestGetGuildInfo) *ResponseGetGuildInfo
}

// New 构造函数
func (r *RequestGetGuildInfo) New() any {
	return &RequestGetGuildInfo{
		Request: &Request{
			Action: "get_guild_info",
		},
	}
}

// Do 执行
func (r *RequestGetGuildInfo) Do() (e any) {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetGuildList 获取群组列表
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_guild_list
type RequestGetGuildList struct {
	*Request
	// f
	f func(events *RequestGetGuildList) *ResponseGetGuildList
}

// New 构造函数
func (*RequestGetGuildList) New() any {
	return &RequestGetGuildList{
		Request: &Request{
			Action: "get_guild_list",
		},
	}
}

// Do 执行
func (r *RequestGetGuildList) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestSetGuildName 设置群组名称
//
// Reference: https://12.onebot.dev/interface/guild/actions/#set_guild_name
type RequestSetGuildName struct {
	*Request
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// GuildName 新群组名称
	GuildName string `json:"guild_name"`
	// f
	f func(events *RequestSetGuildName) *ResponseSetGuildName
}

// New 构造函数
func (*RequestSetGuildName) New() any {
	return &RequestSetGuildName{
		Request: &Request{
			Action: "set_guild_name",
		},
	}
}

// Do 执行
func (r *RequestSetGuildName) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetGuildMemberInfo 获取群组成员信息
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_guild_member_info
type RequestGetGuildMemberInfo struct {
	*Request
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// UserId 用户 ID
	UserId string `json:"user_id"`
	//f
	f func(events *RequestGetGuildMemberInfo) *ResponseGetGuildMemberInfo
}

// New 构造函数
func (*RequestGetGuildMemberInfo) New() any {
	return &RequestGetGuildMemberInfo{
		Request: &Request{
			Action: "get_guild_member_info",
		},
	}
}

// Do 执行
func (r *RequestGetGuildMemberInfo) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetGuildMemberList 获取群组成员列表
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_guild_member_list
type RequestGetGuildMemberList struct {
	*Request
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// f
	f func(events *RequestGetGuildMemberList) *ResponseGetGuildMemberList
}

// New 构造函数
func (*RequestGetGuildMemberList) New() any {
	return &RequestGetGuildMemberList{
		Request: &Request{
			Action: "get_guild_member_list",
		},
	}
}

// Do 执行
func (r *RequestGetGuildMemberList) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestLeaveGuild 退出群组
//
// Reference: https://12.onebot.dev/interface/guild/actions/#leave_guild
type RequestLeaveGuild struct {
	*Request
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// f
	f func(events *RequestLeaveGuild) *ResponseLeaveGuild
}

// New 构造函数
func (*RequestLeaveGuild) New() any {
	return &RequestLeaveGuild{
		Request: &Request{
			Action: "leave_guild",
		},
	}
}

// Do 执行
func (r *RequestLeaveGuild) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetChannelInfo 获取频道信息
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_channel_info
type RequestGetChannelInfo struct {
	*Request
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// ChannelId 频道 ID
	ChannelId string `json:"channel_id"`
	// f
	f func(events *RequestGetChannelInfo) *ResponseGetChannelInfo
}

// New 构造函数
func (*RequestGetChannelInfo) New() any {
	return &RequestGetChannelInfo{
		Request: &Request{
			Action: "get_channel_info",
		},
	}
}

// Do 执行
func (r *RequestGetChannelInfo) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetChannelList 获取频道列表
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_channel_list
type RequestGetChannelList struct {
	*Request
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// JoinedOnly 只获取机器人账号已加入的频道列表
	JoinedOnly bool `json:"joined_only"`
	// f
	f func(events *RequestGetChannelList) *ResponseGetChannelList
}

// New 构造函数
func (*RequestGetChannelList) New() any {
	return &RequestGetChannelList{
		Request: &Request{
			Action: "get_channel_list",
		},
		JoinedOnly: false,
	}
}

// Do 执行
func (r *RequestGetChannelList) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestSetChannelName 设置频道名称
//
// Reference: https://12.onebot.dev/interface/guild/actions/#set_channel_name
type RequestSetChannelName struct {
	*Request
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// ChannelId 频道 ID
	ChannelId string `json:"channel_id"`
	// ChannelName 新频道名称
	ChannelName string `json:"channel_name"`
	// f
	f func(events *RequestSetChannelName) *ResponseSetChannelName
}

// New 构造函数
func (*RequestSetChannelName) New() any {
	return &RequestSetChannelName{
		Request: &Request{
			Action: "set_channel_name",
		},
	}
}

// Do 执行
func (r *RequestSetChannelName) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetChannelMemberInfo 获取频道成员信息
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_channel_member_info
type RequestGetChannelMemberInfo struct {
	*Request
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// ChannelId 频道 ID
	ChannelId string `json:"channel_id"`
	// UserId 用户 ID
	UserId string `json:"user_id"`
	// f
	f func(events *RequestGetChannelMemberInfo) *ResponseGetChannelMemberInfo
}

// New 构造函数
func (*RequestGetChannelMemberInfo) New() any {
	return &RequestGetChannelMemberInfo{
		Request: &Request{
			Action: "get_channel_member_info",
		},
	}
}

// Do 执行
func (r *RequestGetChannelMemberInfo) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetChannelMemberList 获取频道成员列表
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_channel_member_list
type RequestGetChannelMemberList struct {
	*Request
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// ChannelId 频道 ID
	ChannelId string `json:"channel_id"`
	// f
	f func(events *RequestGetChannelMemberList) *ResponseGetChannelMemberList
}

// New 构造函数
func (*RequestGetChannelMemberList) New() any {
	return &RequestGetChannelMemberList{
		Request: &Request{
			Action: "get_channel_member_list",
		},
	}
}

// Do 执行
func (r *RequestGetChannelMemberList) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestLeaveChannel 离开频道
//
// Reference: https://12.onebot.dev/interface/guild/actions/#leave_channel
type RequestLeaveChannel struct {
	*Request
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// ChannelId 频道 ID
	ChannelId string `json:"channel_id"`
	// f
	f func(events *RequestLeaveChannel) *ResponseLeaveChannel
}

// New 构造函数
func (*RequestLeaveChannel) New() any {
	return &RequestLeaveChannel{
		Request: &Request{
			Action: "leave_channel",
		},
	}
}

// Do 执行
func (r *RequestLeaveChannel) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestUploadFile 上传文件
//
// Reference: https://12.onebot.dev/interface/file/actions/#upload_file
type RequestUploadFile struct {
	*Request
	// Name 文件名，如 foo.jpg
	Name string `json:"name"`
	// Url 文件 URL，当 type 为 url 时必须返回，应用端必须能以 HTTP(S) 协议从此 URL 下载文件
	Url string `json:"url"`
	// Headers	下载 URL 时需要添加的 HTTP 请求头，可选返回
	Headers map[string]string `json:"headers"`
	// Path 文件路径，当 type 为 path 时必须返回，应用端必须能从此路径访问文件
	Path string `json:"path"`
	// Data 文件数据，当 type 为 data 时必须返回
	Data []byte `json:"data"`
	// Sha256 文件数据（原始二进制）的 SHA256 校验和，全小写，可选返回
	Sha256 string `json:"sha256"`
	// Type 	上传文件的方式，可以为 url、path、data 或扩展的方式
	Type string `json:"type"`
	// f
	f func(events *RequestUploadFile) *ResponseUploadFile
}

// New 构造函数
func (*RequestUploadFile) New() any {
	return &RequestUploadFile{
		Request: &Request{
			Action: "upload_file",
		},
		Headers: make(map[string]string),
	}
}

// Do 执行
func (r *RequestUploadFile) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestUploadFileFragmented 上传文件分片
//
// Reference: https://12.onebot.dev/interface/file/actions/#upload_file_fragmented
type RequestUploadFileFragmented struct {
	*Request
	// Stage 上传阶段，必须为 prepare 传输阶段，必须为 transfer
	Stage string `json:"stage"`
	// Name 文件名，如 foo.jpg
	Name string `json:"name"`
	// TotalSize 文件总大小，单位为字节
	TotalSize int64 `json:"total_size"`
	// FileID 准备阶段返回的文件 ID
	FileID string `json:"file_id"`
	// Offset 	本次传输的文件偏移，单位：字节
	Offset int64 `json:"offset"`
	// Data 本次传输的文件数据
	Data []byte `json:"data"`
	// Sha256 本次传输的文件数据（原始二进制）的 SHA256 校验和，全小写
	Sha256 string `json:"sha256"`
	// f
	f func(events *RequestUploadFileFragmented) *ResponseUploadFileFragmented
}

// New 构造函数
func (*RequestUploadFileFragmented) New() any {
	return &RequestUploadFileFragmented{
		Request: &Request{
			Action: "upload_file_fragmented",
		},
	}
}

// Do 执行
func (r *RequestUploadFileFragmented) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetFile 获取文件
//
// Reference: https://12.onebot.dev/interface/file/actions/#get_file
type RequestGetFile struct {
	*Request
	// FileId 文件 ID
	FileId string `json:"file_id"`
	// Type 获取文件的方式，可以为 url、path、data 或扩展的方式
	Type string `json:"type"`
	// f
	f func(events *RequestGetFile) *ResponseGetFile
}

// New 构造函数
func (*RequestGetFile) New() any {
	return &RequestGetFile{
		Request: &Request{
			Action: "get_file",
		},
	}
}

// Do 执行
func (r *RequestGetFile) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}

// RequestGetFileFragmented 获取文件分片
//
// Reference: https://12.onebot.dev/interface/file/actions/#get_file_fragmented
type RequestGetFileFragmented struct {
	*Request
	// Stage 上传阶段，必须为 prepare 传输阶段，必须为 transfer
	Stage string `json:"stage"`
	// FileID 文件 ID
	FileID string `json:"file_id"`
	// Offset 本次传输的文件偏移，单位：字节
	Offset int64 `json:"offset"`
	// Size 本次传输的文件大小，单位：字节
	Size int64 `json:"size"`
	// f
	f func(events *RequestGetFileFragmented) *ResponseGetFileFragmented
}

// New 构造函数
func (*RequestGetFileFragmented) New() any {
	return &RequestGetFileFragmented{
		Request: &Request{
			Action: "get_file_fragmented",
		},
	}
}

// Do 执行
func (r *RequestGetFileFragmented) Do() any {
	if r.f == nil {
		return NewEmptyResponse(ResponseCodeBadHandler)
	}
	return r.f(r)
}
