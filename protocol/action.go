package protocol

/*

	Action: 动作
	Reference: https://12.onebot.dev/glossary/#event
	分为Request和Response
	Request为sdk向实现端发送的请求, Response为实现端收到请求后返回的响应
	这里把OneBot12标准的所有动作都实现了, 可以通过HandleActionXXX函数来注册自定义的动作
	该函数会返回一个action名称和一个RequestInterface接口
	如果想自定义动作, 可以通过实现RequestInterface接口来实现

*/

/*
	另外, 本代码大量由copilot生成, 生成错的地方请打骂copilot
*/

// HandleActionGetLatestEvents 处理GetLatestEvents请求
func HandleActionGetLatestEvents(f func(events *RequestGetLatestEvents) *ResponseGetLatestEvents) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetLatestEvents{}
	r := a.New()
	r.(*RequestGetLatestEvents).f = f
	return r.(*RequestGetLatestEvents)
}

// HandleActionGetSupportedActions 处理GetSupportedActions请求
func HandleActionGetSupportedActions(f func(events *RequestGetSupportedActions) *ResponseGetSupportedActions) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetSupportedActions{}
	r := a.New()
	r.(*RequestGetSupportedActions).f = f
	return r.(*RequestGetSupportedActions)
}

// HandleActionGetStatus 处理GetStatus请求
func HandleActionGetStatus(f func(events *RequestGetStatus) *ResponseGetStatus) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetStatus{}
	r := a.New()
	r.(*RequestGetStatus).f = f
	return r.(*RequestGetStatus)
}

// HandleActionGetVersion 处理GetVersion请求
func HandleActionGetVersion(f func(events *RequestGetVersion) *ResponseGetVersion) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetVersion{}
	r := a.New()
	r.(*RequestGetVersion).f = f
	return r.(*RequestGetVersion)
}

// HandleActionSendMessage 处理SendMessage请求
func HandleActionSendMessage(f func(events *RequestSendMessage) *ResponseSendMessage) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestSendMessage{}
	r := a.New()
	r.(*RequestSendMessage).f = f
	return r.(*RequestSendMessage)
}

// HandleActionDeleteMessage 处理DeleteMessage请求
func HandleActionDeleteMessage(f func(events *RequestDeleteMessage) *ResponseDeleteMessage) (string, RequestInterface) {
	if f == nil {
		return "", nil
	}
	a := RequestDeleteMessage{}
	r := a.New()
	r.(*RequestDeleteMessage).f = f
	return r.(*RequestDeleteMessage).Action, r.(*RequestDeleteMessage)
}

// HandleActionGetSelfInfo 处理GetSelfInfo请求
func HandleActionGetSelfInfo(f func(events *RequestGetSelfInfo) *ResponseGetSelfInfo) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetSelfInfo{}
	r := a.New()
	r.(*RequestGetSelfInfo).f = f
	return r.(*RequestGetSelfInfo)
}

// HandleActionGetUserInfo 处理GetUserInfo请求
func HandleActionGetUserInfo(f func(events *RequestGetUserInfo) *ResponseGetUserInfo) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetUserInfo{}
	r := a.New()
	r.(*RequestGetUserInfo).f = f
	return r.(*RequestGetUserInfo)
}

// HandleActionGetFriendList 处理GetFriendList请求
func HandleActionGetFriendList(f func(events *RequestGetFriendList) *ResponseGetFriendList) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetFriendList{}
	r := a.New()
	r.(*RequestGetFriendList).f = f
	return r.(*RequestGetFriendList)
}

// HandleActionGetGroupInfo 处理GetGroupInfo请求
func HandleActionGetGroupInfo(f func(events *RequestGetGroupInfo) *ResponseGetGroupInfo) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetGroupInfo{}
	r := a.New()
	r.(*RequestGetGroupInfo).f = f
	return r.(*RequestGetGroupInfo)
}

// HandleActionGetGroupList 处理GetGroupList请求
func HandleActionGetGroupList(f func(events *RequestGetGroupList) *ResponseGetGroupList) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetGroupList{}
	r := a.New()
	r.(*RequestGetGroupList).f = f
	return r.(*RequestGetGroupList)
}

// HandleActionGetGroupMemberInfo 处理GetGroupMemberInfo请求
func HandleActionGetGroupMemberInfo(f func(events *RequestGetGroupMemberInfo) *ResponseGetGroupMemberInfo) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetGroupMemberInfo{}
	r := a.New()
	r.(*RequestGetGroupMemberInfo).f = f
	return r.(*RequestGetGroupMemberInfo)
}

// HandleActionGetGroupMemberList 处理GetGroupMemberList请求
func HandleActionGetGroupMemberList(f func(events *RequestGetGroupMemberList) *ResponseGetGroupMemberList) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetGroupMemberList{}
	r := a.New()
	r.(*RequestGetGroupMemberList).f = f
	return r.(*RequestGetGroupMemberList)
}

// HandleActionSetGroupName 处理SetGroupName请求
func HandleActionSetGroupName(f func(events *RequestSetGroupName) *ResponseSetGroupName) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestSetGroupName{}
	r := a.New()
	r.(*RequestSetGroupName).f = f
	return r.(*RequestSetGroupName)
}

// HandleActionLeaveGroup 处理LeaveGroup请求
func HandleActionLeaveGroup(f func(events *RequestLeaveGroup) *ResponseLeaveGroup) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestLeaveGroup{}
	r := a.New()
	r.(*RequestLeaveGroup).f = f
	return r.(*RequestLeaveGroup)
}

// HandleActionGetGuildInfo 处理GetGuildInfo请求
func HandleActionGetGuildInfo(f func(events *RequestGetGuildInfo) *ResponseGetGuildInfo) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetGuildInfo{}
	r := a.New()
	r.(*RequestGetGuildInfo).f = f
	return r.(*RequestGetGuildInfo)
}

// HandleActionGetGuildList 处理GetGuildList请求
func HandleActionGetGuildList(f func(events *RequestGetGuildList) *ResponseGetGuildList) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetGuildList{}
	r := a.New()
	r.(*RequestGetGuildList).f = f
	return r.(*RequestGetGuildList)
}

// HandleActionSetGuildName 处理SetGuildName请求
func HandleActionSetGuildName(f func(events *RequestSetGuildName) *ResponseSetGuildName) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestSetGuildName{}
	r := a.New()
	r.(*RequestSetGuildName).f = f
	return r.(*RequestSetGuildName)
}

// HandleActionGetGuildMemberInfo 处理GetGuildMemberInfo请求
func HandleActionGetGuildMemberInfo(f func(events *RequestGetGuildMemberInfo) *ResponseGetGuildMemberInfo) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetGuildMemberInfo{}
	r := a.New()
	r.(*RequestGetGuildMemberInfo).f = f
	return r.(*RequestGetGuildMemberInfo)
}

// HandleActionGetGuildMemberList 处理GetGuildMemberList请求
func HandleActionGetGuildMemberList(f func(events *RequestGetGuildMemberList) *ResponseGetGuildMemberList) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetGuildMemberList{}
	r := a.New()
	r.(*RequestGetGuildMemberList).f = f
	return r.(*RequestGetGuildMemberList)
}

// HandleActionLeaveGuild 处理LeaveGuild请求
func HandleActionLeaveGuild(f func(events *RequestLeaveGuild) *ResponseLeaveGuild) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestLeaveGuild{}
	r := a.New()
	r.(*RequestLeaveGuild).f = f
	return r.(*RequestLeaveGuild)
}

// HandleActionGetChannelInfo 处理GetChannelInfo请求
func HandleActionGetChannelInfo(f func(events *RequestGetChannelInfo) *ResponseGetChannelInfo) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetChannelInfo{}
	r := a.New()
	r.(*RequestGetChannelInfo).f = f
	return r.(*RequestGetChannelInfo)
}

// HandleActionGetChannelList 处理GetChannelList请求
func HandleActionGetChannelList(f func(events *RequestGetChannelList) *ResponseGetChannelList) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetChannelList{}
	r := a.New()
	r.(*RequestGetChannelList).f = f
	return r.(*RequestGetChannelList)
}

// HandleActionSetChannelName 处理SetChannelName请求
func HandleActionSetChannelName(f func(events *RequestSetChannelName) *ResponseSetChannelName) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestSetChannelName{}
	r := a.New()
	r.(*RequestSetChannelName).f = f
	return r.(*RequestSetChannelName)
}

// HandleActionGetChannelMemberInfo 处理GetChannelMemberInfo请求
func HandleActionGetChannelMemberInfo(f func(events *RequestGetChannelMemberInfo) *ResponseGetChannelMemberInfo) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetChannelMemberInfo{}
	r := a.New()
	r.(*RequestGetChannelMemberInfo).f = f
	return r.(*RequestGetChannelMemberInfo)
}

// HandleActionGetChannelMemberList 处理GetChannelMemberList请求
func HandleActionGetChannelMemberList(f func(events *RequestGetChannelMemberList) *ResponseGetChannelMemberList) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetChannelMemberList{}
	r := a.New()
	r.(*RequestGetChannelMemberList).f = f
	return r.(*RequestGetChannelMemberList)
}

// HandleActionLeaveChannel 处理LeaveChannel请求
func HandleActionLeaveChannel(f func(events *RequestLeaveChannel) *ResponseLeaveChannel) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestLeaveChannel{}
	r := a.New()
	r.(*RequestLeaveChannel).f = f
	return r.(*RequestLeaveChannel)
}

// HandleActionUploadFile 处理UploadFile请求
func HandleActionUploadFile(f func(events *RequestUploadFile) *ResponseUploadFile) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestUploadFile{}
	r := a.New()
	r.(*RequestUploadFile).f = f
	return r.(*RequestUploadFile)
}

// HandleActionUploadFileFragmented 处理UploadFileFragmented请求
func HandleActionUploadFileFragmented(f func(events *RequestUploadFileFragmented) *ResponseUploadFileFragmented) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestUploadFileFragmented{}
	r := a.New()
	r.(*RequestUploadFileFragmented).f = f
	return r.(*RequestUploadFileFragmented)
}

// HandleActionGetFile 处理GetFile请求
func HandleActionGetFile(f func(events *RequestGetFile) *ResponseGetFile) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetFile{}
	r := a.New()
	r.(*RequestGetFile).f = f
	return r.(*RequestGetFile)
}

// HandleActionGetFileFragmented 处理GetFileFragmented请求
func HandleActionGetFileFragmented(f func(events *RequestGetFileFragmented) *ResponseGetFileFragmented) RequestInterface {
	if f == nil {
		return nil
	}
	a := RequestGetFileFragmented{}
	r := a.New()
	r.(*RequestGetFileFragmented).f = f
	return r.(*RequestGetFileFragmented)
}
