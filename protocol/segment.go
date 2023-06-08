package protocol

// Segment 消息段
//
// Reference: https://12.onebot.dev/glossary/#message-segment-segment
type Segment struct {
	// Type 消息段类型
	Type string `json:"type"`
	// Data 消息段数据
	Data map[string]interface{} `json:"data"`
}

// GetSegment 返回消息段
//
// Reference: https://12.onebot.dev/glossary/#message-segment-segment
func GetSegment(t string, data map[string]interface{}) Segment {
	return Segment{
		Type: t,
		Data: data,
	}
}

// GetSegmentText 返回文本消息段
//
// Reference: https://12.onebot.dev/interface/message/segments/#text
func GetSegmentText(text string) Segment {
	return GetSegment("text", map[string]interface{}{
		"text": text,
	})
}

// GetSegmentMention 返回@消息段
//
// Reference: https://12.onebot.dev/interface/message/segments/#mention
func GetSegmentMention(userId string) Segment {
	return GetSegment("mention", map[string]interface{}{
		"user_id": userId,
	})
}

// GetSegmentMentionAll 返回@全体成员消息段
//
// Reference: https://12.onebot.dev/interface/message/segments/#mention_all
func GetSegmentMentionAll() Segment {
	return GetSegment("mention_all", map[string]interface{}{})
}

// GetSegmentImage 返回图片消息段
//
// Reference: https://12.onebot.dev/interface/message/segments/#image
func GetSegmentImage(fileId string) Segment {
	return GetSegment("image", map[string]interface{}{
		"file_id": fileId,
	})
}

// GetSegmentVoice 返回语音消息段
//
// Reference: https://12.onebot.dev/interface/message/segments/#voice
func GetSegmentVoice(fileId string) Segment {
	return GetSegment("voice", map[string]interface{}{
		"file_id": fileId,
	})
}

// GetSegmentAudio 返回音频消息段
//
// Reference: https://12.onebot.dev/interface/message/segments/#audio
func GetSegmentAudio(fileId string) Segment {
	return GetSegment("audio", map[string]interface{}{
		"file_id": fileId,
	})
}

// GetSegmentVideo 返回视频消息段
//
// Reference: https://12.onebot.dev/interface/message/segments/#video
func GetSegmentVideo(fileId string) Segment {
	return GetSegment("video", map[string]interface{}{
		"file_id": fileId,
	})
}

// GetSegmentFile 返回文件消息段
//
// Reference: https://12.onebot.dev/interface/message/segments/#file
func GetSegmentFile(fileId string) Segment {
	return GetSegment("file", map[string]interface{}{
		"file_id": fileId,
	})
}

// GetSegmentLocation 返回位置消息段
//
// Reference: https://12.onebot.dev/interface/message/segments/#location
func GetSegmentLocation(latitude float64, longitude float64, title string, content string) Segment {
	return GetSegment("location", map[string]interface{}{
		"latitude":  latitude,
		"longitude": longitude,
		"title":     title,
		"content":   content,
	})
}

// GetSegmentReply 返回回复消息段
//
// Reference: https://12.onebot.dev/interface/message/segments/#reply
func GetSegmentReply(messageId string, UserId string) Segment {
	return GetSegment("reply", map[string]interface{}{
		"message_id": messageId,
		"user_id":    UserId,
	})
}
