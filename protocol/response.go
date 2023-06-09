package protocol

import "github.com/FishZe/go-libonebot/util"

// EmptyResponse 空响应
type EmptyResponse struct {
	*Response
}

// NewEmptyResponse 构造函数
func NewEmptyResponse(retCode int) *EmptyResponse {
	return &EmptyResponse{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetLatestEvents 获取最近的事件
type ResponseGetLatestEvents struct {
	*Response
	// Events 事件列表
	Events []any
}

// NewResponseGetLatestEvents 获取最近的事件 构造函数
func NewResponseGetLatestEvents(retCode int) *ResponseGetLatestEvents {
	return &ResponseGetLatestEvents{
		Response: &Response{
			Retcode: retCode,
		},
		Events: make([]any, 0),
	}
}

// ResponseGetSupportedActions 获取支持的动作列表
//
// Reference: https://12.onebot.dev/interface/meta/actions/#get_supported_actions
type ResponseGetSupportedActions struct {
	*Response
	// Actions 动作列表
	Actions []string
}

// NewResponseGetSupportedActions 获取支持的动作列表 构造函数
func NewResponseGetSupportedActions(retCode int) *ResponseGetSupportedActions {
	return &ResponseGetSupportedActions{
		Response: &Response{
			Retcode: retCode,
		},
		Actions: make([]string, 0),
	}
}

// ResponseGetStatus 获取运行状态
//
// Reference: https://12.onebot.dev/interface/meta/actions/#get_status
type ResponseGetStatus struct {
	*Response

	Good bool `json:"good"`
	Bots []struct {
		Self   Self `json:"self"`
		Online bool `json:"online"`
		Extra  map[string]any
	} `json:"bots"`
}

// NewResponseGetStatus 获取运行状态 构造函数
func NewResponseGetStatus(retCode int) *ResponseGetStatus {
	return &ResponseGetStatus{
		Response: &Response{
			Retcode: retCode,
		},
		Bots: make([]struct {
			Self   Self `json:"self"`
			Online bool `json:"online"`
			Extra  map[string]any
		}, 0),
	}
}

// ResponseGetVersion 获取版本信息
//
// Reference: https://12.onebot.dev/interface/meta/actions/#get_version
type ResponseGetVersion struct {
	*Response
	// Impl OneBot 实现名称，格式见 术语表
	Impl string `json:"impl"`
	// Version OneBot 实现的版本号
	Version string `json:"version"`
	// OneBotVersion OneBot 实现的 OneBot 标准版本号
	OnebotVersion string `json:"onebot_version"`
}

// NewResponseGetVersion 获取版本信息 构造函数
func NewResponseGetVersion(retCode int) *ResponseGetVersion {
	return &ResponseGetVersion{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseSendMessage 发送消息
//
// Reference: https://12.onebot.dev/interface/message/actions/#send_message
type ResponseSendMessage struct {
	*Response
	// MessageId 消息 ID
	MessageId string `json:"message_id"`
	// Time 消息成功发出的时间（Unix 时间戳），单位：秒
	Time float64 `json:"time"`
}

// NewResponseSendMessage 发送消息 构造函数
//
// 已经帮你生成好了MessageID和Time 直接返回就行
func NewResponseSendMessage(retCode int) *ResponseSendMessage {
	if retCode == 0 {
		return &ResponseSendMessage{
			Response: &Response{
				Retcode: retCode,
			},
			MessageId: util.GetUUID(),
			Time:      util.GetTimeStampFloat64(),
		}
	} else {
		return &ResponseSendMessage{
			Response: &Response{
				Retcode: retCode,
			},
			Time: util.GetTimeStampFloat64(),
		}
	}
}

// ResponseDeleteMessage 撤回消息
//
// Reference: https://12.onebot.dev/interface/message/actions/#delete_message
//
// 空的
type ResponseDeleteMessage struct {
	*Response
}

// NewResponseDeleteMessage 撤回消息 构造函数
func NewResponseDeleteMessage(retCode int) *ResponseDeleteMessage {
	return &ResponseDeleteMessage{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetSelfInfo 获取机器人自身信息
//
// Reference: https://12.onebot.dev/interface/user/actions/#get_self_info
type ResponseGetSelfInfo struct {
	*Response
	// UserId 机器人用户 ID
	UserId string `json:"user_id"`
	// Nickname 机器人名称/姓名/昵称
	UserName string `json:"user_name"`
	// UserDisPlayName 机器人账号设置的显示名称，若无则为空字符串
	UserDisplayName string `json:"user_displayname"`
}

// NewResponseGetSelfInfo 获取机器人自身信息
func NewResponseGetSelfInfo(retCode int) *ResponseGetSelfInfo {
	return &ResponseGetSelfInfo{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetUserInfo 获取用户信息
//
// Reference: https://12.onebot.dev/interface/user/actions/#get_user_info
type ResponseGetUserInfo struct {
	*Response
	// UserId 用户 ID
	UserId string `json:"user_id"`
	// UserName 用户名称/姓名/昵称
	UserName string `json:"user_name"`
	// UserDisplayName 用户设置的显示名称，若无则为空字符串
	UserDisplayName string `json:"user_displayname"`
	// UserRemark 机器人账号对该用户的备注名称，若无则为空字符串
	UserRemark string `json:"user_remark"`
}

// NewResponseGetUserInfo 获取用户信息 构造函数
func NewResponseGetUserInfo(retCode int) *ResponseGetUserInfo {
	return &ResponseGetUserInfo{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetFriendList 获取好友列表
//
// Reference: https://12.onebot.dev/interface/user/actions/#get_friend_list
type ResponseGetFriendList struct {
	*Response
	// Friends 好友列表
	Friends []struct {
		// UserId 用户 ID
		UserId string `json:"user_id"`
		// UserName 用户名称/姓名/昵称
		UserName string `json:"user_name"`
		// UserDisplayName 用户设置的显示名称，若无则为空字符串
		UserDisplayName string `json:"user_displayname"`
		// UserRemark 机器人账号对该用户的备注名称，若无则为空字符串
		UserRemark string `json:"user_remark"`
	}
}

// NewResponseGetFriendList 获取好友列表 构造函数
func NewResponseGetFriendList(retCode int) *ResponseGetFriendList {
	return &ResponseGetFriendList{
		Response: &Response{
			Retcode: retCode,
		},
		Friends: make([]struct {
			UserId          string `json:"user_id"`
			UserName        string `json:"user_name"`
			UserDisplayName string `json:"user_displayname"`
			UserRemark      string `json:"user_remark"`
		}, 0),
	}
}

// ResponseGetGroupInfo 获取群信息
//
// Reference: https://12.onebot.dev/interface/group/actions/#get_group_info
type ResponseGetGroupInfo struct {
	*Response
	// GroupId 群号
	GroupId string `json:"group_id"`
	// GroupName 群名称
	GroupName string `json:"group_name"`
}

// NewResponseGetGroupInfo 获取群信息 构造函数
func NewResponseGetGroupInfo(retCode int) *ResponseGetGroupInfo {
	return &ResponseGetGroupInfo{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetGroupList 获取群列表
//
// Reference: https://12.onebot.dev/interface/group/actions/#get_group_list
type ResponseGetGroupList struct {
	*Response
	Groups []struct {
		// GroupId 群号
		GroupId string `json:"group_id"`
		// GroupName 群名称
		GroupName string `json:"group_name"`
	}
}

// NewResponseGetGroupList 获取群列表 构造函数
func NewResponseGetGroupList(retCode int) *ResponseGetGroupList {
	return &ResponseGetGroupList{
		Response: &Response{
			Retcode: retCode,
		},
		Groups: make([]struct {
			GroupId   string `json:"group_id"`
			GroupName string `json:"group_name"`
		}, 0),
	}
}

// ResponseGetGroupMemberInfo 获取群成员信息
//
// Reference: https://12.onebot.dev/interface/group/actions/#get_group_member_info
type ResponseGetGroupMemberInfo struct {
	*Response
	// UserId 用户 ID
	UserId string `json:"user_id"`
	// UserName 用户名称/姓名/昵称
	UserName string `json:"user_name"`
	// UserDisplayName 用户设置的群内显示名称或账号显示名称，若无则为空字符串
	UserDisplayName string `json:"user_displayname"`
}

// NewResponseGetGroupMemberInfo 获取群成员信息 构造函数
func NewResponseGetGroupMemberInfo(retCode int) *ResponseGetGroupMemberInfo {
	return &ResponseGetGroupMemberInfo{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetGroupMemberList 获取群成员列表
//
// Reference: https://12.onebot.dev/interface/group/actions/#get_group_member_list
type ResponseGetGroupMemberList struct {
	*Response
	// Members 群成员列表
	Members []struct {
		// UserId 用户 ID
		UserId string `json:"user_id"`
		// UserName 用户名称/姓名/昵称
		UserName string `json:"user_name"`
		// UserDisplayName 用户设置的群内显示名称或账号显示名称，若无则为空字符串
		UserDisplayName string `json:"user_displayname"`
	}
}

// NewResponseGetGroupMemberList 获取群成员列表 构造函数
func NewResponseGetGroupMemberList(retCode int) *ResponseGetGroupMemberList {
	return &ResponseGetGroupMemberList{
		Response: &Response{
			Retcode: retCode,
		},
		Members: make([]struct {
			UserId          string `json:"user_id"`
			UserName        string `json:"user_name"`
			UserDisplayName string `json:"user_displayname"`
		}, 0),
	}
}

// ResponseSetGroupName 设置群名称
//
// Reference: https://12.onebot.dev/interface/group/actions/#set_group_name
//
// 空的
type ResponseSetGroupName struct {
	*Response
}

// NewResponseSetGroupName 设置群名称 构造函数
func NewResponseSetGroupName(retCode int) *ResponseSetGroupName {
	return &ResponseSetGroupName{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseLeaveGroup 退群
//
// Reference: https://12.onebot.dev/interface/group/actions/#leave_group
//
// 空的
type ResponseLeaveGroup struct {
	*Response
}

// NewResponseLeaveGroup 退群 构造函数
func NewResponseLeaveGroup(retCode int) *ResponseLeaveGroup {
	return &ResponseLeaveGroup{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetGuildInfo 获取群组信
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_guild_info
type ResponseGetGuildInfo struct {
	*Response
	// GuildId 群组 ID
	GuildId string `json:"guild_id"`
	// GuildName 群组名称
	GuildName string `json:"guild_name"`
}

// NewResponseGetGuildInfo 获取群组信 构造函数
func NewResponseGetGuildInfo(retCode int) *ResponseGetGuildInfo {
	return &ResponseGetGuildInfo{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetGuildList 获取群组列表
//
// 获取机器人加入的群组列表。
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_guild_list
type ResponseGetGuildList struct {
	*Response
	// Guilds 群组列表
	Guilds []struct {
		// GuildId 群组 ID
		GuildId string `json:"guild_id"`
		// GuildName 群组名称
		GuildName string `json:"guild_name"`
	}
}

// NewResponseGetGuildList 获取群组列表 构造函数
func NewResponseGetGuildList(retCode int) *ResponseGetGuildList {
	return &ResponseGetGuildList{
		Response: &Response{
			Retcode: retCode,
		},
		Guilds: make([]struct {
			GuildId   string `json:"guild_id"`
			GuildName string `json:"guild_name"`
		}, 0),
	}
}

// ResponseSetGuildName 设置群组名称
//
// Reference: https://12.onebot.dev/interface/guild/actions/#set_guild_name
type ResponseSetGuildName struct {
	*Response
	// 空的
}

// NewResponseSetGuildName 设置群组名称 构造函数
func NewResponseSetGuildName(retCode int) *ResponseSetGuildName {
	return &ResponseSetGuildName{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetGuildMemberInfo 获取群组成员信息
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_guild_member_info
type ResponseGetGuildMemberInfo struct {
	*Response
	// UserId 用户 ID
	UserId string `json:"user_id"`
	// UserName 用户名称/姓名/昵称
	UserName string `json:"user_name"`
	// UserDisplayName 	用户设置的群组内显示名称或账号显示名称，若无则为空字符串
	UserDisplayName string `json:"user_displayname"`
}

// NewResponseGetGuildMemberInfo 获取群组成员信息 构造函数
func NewResponseGetGuildMemberInfo(retCode int) *ResponseGetGuildMemberInfo {
	return &ResponseGetGuildMemberInfo{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetGuildMemberList 获取群组成员列表
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_guild_member_list
type ResponseGetGuildMemberList struct {
	*Response
	// Members 群组成员列表
	Members []struct {
		// UserId 用户 ID
		UserId string `json:"user_id"`
		// UserName 用户名称/姓名/昵称
		UserName string `json:"user_name"`
		// UserDisplayName 	用户设置的群组内显示名称或账号显示名称，若无则为空字符串
		UserDisplayName string `json:"user_displayname"`
	}
}

// NewResponseGetGuildMemberList 获取群组成员列表 构造函数
func NewResponseGetGuildMemberList(retCode int) *ResponseGetGuildMemberList {
	return &ResponseGetGuildMemberList{
		Response: &Response{
			Retcode: retCode,
		},
		Members: make([]struct {
			UserId          string `json:"user_id"`
			UserName        string `json:"user_name"`
			UserDisplayName string `json:"user_displayname"`
		}, 0),
	}
}

// ResponseLeaveGuild 退群组
//
// Reference: https://12.onebot.dev/interface/guild/actions/#leave_guild
type ResponseLeaveGuild struct {
	*Response
	// 空的
}

// NewResponseLeaveGuild 退群组 构造函数
func NewResponseLeaveGuild(retCode int) *ResponseLeaveGuild {
	return &ResponseLeaveGuild{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetChannelInfo 获取频道信息
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_channel_info
type ResponseGetChannelInfo struct {
	*Response
	// ChannelId 频道 ID
	ChannelId string `json:"channel_id"`
	// ChannelName 频道名称
	ChannelName string `json:"channel_name"`
}

// NewResponseGetChannelInfo 获取频道信息 构造函数
func NewResponseGetChannelInfo(retCode int) *ResponseGetChannelInfo {
	return &ResponseGetChannelInfo{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetChannelList 获取频道列表
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_channel_list
type ResponseGetChannelList struct {
	*Response
	// Channels 频道列表
	Channels []struct {
		// ChannelId 频道 ID
		ChannelId string `json:"channel_id"`
		// ChannelName 频道名称
		ChannelName string `json:"channel_name"`
	}
}

// NewResponseGetChannelList 获取频道列表 构造函数
func NewResponseGetChannelList(retCode int) *ResponseGetChannelList {
	return &ResponseGetChannelList{
		Response: &Response{
			Retcode: retCode,
		},
		Channels: make([]struct {
			ChannelId   string `json:"channel_id"`
			ChannelName string `json:"channel_name"`
		}, 0),
	}
}

// ResponseSetChannelName 设置频道名称
//
// Reference: https://12.onebot.dev/interface/guild/actions/#set_channel_name
type ResponseSetChannelName struct {
	*Response
	// 空的
}

// NewResponseSetChannelName 设置频道名称 构造函数
func NewResponseSetChannelName(retCode int) *ResponseSetChannelName {
	return &ResponseSetChannelName{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetChannelMemberInfo 获取频道成员信息
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_channel_member_info
type ResponseGetChannelMemberInfo struct {
	*Response
	// UserId 用户 ID
	UserId string `json:"user_id"`
	// UserName 用户名称/姓名/昵称
	UserName string `json:"user_name"`
	// UserDisplayName 	用户设置的群组内显示名称或账号显示名称，若无则为空字符串
	UserDisplayName string `json:"user_displayname"`
}

// NewResponseGetChannelMemberInfo 获取频道成员信息 构造函数
func NewResponseGetChannelMemberInfo(retCode int) *ResponseGetChannelMemberInfo {
	return &ResponseGetChannelMemberInfo{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetChannelMemberList 获取频道成员列表
//
// Reference: https://12.onebot.dev/interface/guild/actions/#get_channel_member_list
type ResponseGetChannelMemberList struct {
	*Response
	// Members 频道成员列表
	Members []struct {
		// UserId 用户 ID
		UserId string `json:"user_id"`
		// UserName 用户名称/姓名/昵称
		UserName string `json:"user_name"`
		// UserDisplayName 	用户设置的群组内显示名称或账号显示名称，若无则为空字符串
		UserDisplayName string `json:"user_displayname"`
	}
}

// NewResponseGetChannelMemberList 获取频道成员列表 构造函数
func NewResponseGetChannelMemberList(retCode int) *ResponseGetChannelMemberList {
	return &ResponseGetChannelMemberList{
		Response: &Response{
			Retcode: retCode,
		},
		Members: make([]struct {
			UserId          string `json:"user_id"`
			UserName        string `json:"user_name"`
			UserDisplayName string `json:"user_displayname"`
		}, 0),
	}
}

// ResponseLeaveChannel 退出频道
//
// Reference: https://12.onebot.dev/interface/guild/actions/#leave_channel
type ResponseLeaveChannel struct {
	*Response
	// 空的
}

// NewResponseLeaveChannel 退出频道 构造函数
func NewResponseLeaveChannel(retCode int) *ResponseLeaveChannel {
	return &ResponseLeaveChannel{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseUploadFile 上传文件
//
// Reference: https://12.onebot.dev/interface/file/actions/#upload_file
type ResponseUploadFile struct {
	*Response
	// FileId 文件 ID
	FileId string `json:"file_id"`
}

// NewResponseUploadFile 上传文件 构造函数
func NewResponseUploadFile(retCode int) *ResponseUploadFile {
	return &ResponseUploadFile{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseUploadFileFragmented 上传分片文件
//
// Reference: https://12.onebot.dev/interface/file/actions/#upload_file_fragmented
type ResponseUploadFileFragmented struct {
	*Response
	// FileId 文件 ID，仅传输阶段使用
	FileId string `json:"file_id"`
}

// NewResponseUploadFileFragmentedStart 上传分片文件 构造函数
func NewResponseUploadFileFragmentedStart(retCode int) *ResponseUploadFileFragmented {
	return &ResponseUploadFileFragmented{
		Response: &Response{
			Retcode: retCode,
		},
	}
}

// ResponseGetFile 获取文件
//
// Reference: https://12.onebot.dev/interface/file/actions/#get_file
//
// 这里虽然说“必须返回”，但如果平台真的无法获得 URL，当用户请求 type 为 url 时，可以返回 10004 Unsupported Param。具体见 接口定义 - 概述 中对 OneBot 实现的要求。
type ResponseGetFile struct {
	*Response
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
}

// NewResponseGetFile 获取文件 构造函数
func NewResponseGetFile(retCode int) *ResponseGetFile {
	return &ResponseGetFile{
		Response: &Response{
			Retcode: retCode,
		},
		Headers: make(map[string]string),
	}
}

// ResponseGetFileFragmented 获取分片文件
//
// Reference: https://12.onebot.dev/interface/file/actions/#get_file_fragmented
type ResponseGetFileFragmented struct {
	*Response
	// Name 文件名，如 foo.jpg
	Name string `json:"name"`
	// TotalSize 文件完整大小，单位：字节
	TotalSize int `json:"total_size"`
	// Sha256 文件数据（原始二进制）的 SHA256 校验和，全小写
	Sha256 string `json:"sha256"`
	// Data 本次传输的文件数据
	Data []byte `json:"data"`
}

// NewResponseGetFileFragmentedStart 获取分片文件 构造函数
func NewResponseGetFileFragmentedStart(retCode int) *ResponseGetFileFragmented {
	return &ResponseGetFileFragmented{
		Response: &Response{
			Retcode: retCode,
		},
	}
}
