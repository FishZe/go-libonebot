package v11_adapter

import (
	"github.com/FishZe/go-libonebot/protocol"
	"strings"
)

// deleteFromString 从字符串中删除子串
func deleteFromString(s string, t string) string {
	return strings.Replace(s, t, "", -1)
}

// Segment12To11 将CQ码转换为Segment11
func Segment12To11(seg protocol.Segment) protocol.Segment {
	var res = protocol.Segment{
		Data: make(map[string]interface{}),
	}
	switch seg.Type {
	case "mention":
		res.Type = "at"
		if _, ok := seg.Data["user_id"]; ok && seg.Data["user_id"] != "" {
			res.Data["qq"] = seg.Data["user_id"]
		}
	case "mention_all":
		res.Type = "at"
		res.Data["qq"] = "all"
	case "voice":
		res.Type = "record"
		if _, ok := seg.Data["file_id"]; ok && seg.Data["file_id"] != "" {
			res.Data["file"] = seg.Data["file_id"]
		}
	case "video":
		res.Type = "video"
		if _, ok := seg.Data["file_id"]; ok && seg.Data["file_id"] != "" {
			res.Data["file"] = seg.Data["file_id"]
		}
	case "image":
		res.Type = "image"
		if _, ok := seg.Data["file_id"]; ok && seg.Data["file_id"] != "" {
			res.Data["file"] = seg.Data["file_id"]
		}
	case "location":
		res.Type = "location"
		if _, ok := seg.Data["latitude"]; ok && seg.Data["latitude"] != "" {
			res.Data["lat"] = seg.Data["latitude"]
		}
		if _, ok := seg.Data["longitude"]; ok && seg.Data["longitude"] != "" {
			res.Data["lon"] = seg.Data["longitude"]
		}
		if _, ok := seg.Data["title"]; ok && seg.Data["title"] != "" {
			res.Data["title"] = seg.Data["title"]
		}
		if _, ok := seg.Data["content"]; ok && seg.Data["content"] != "" {
			res.Data["content"] = seg.Data["content"]
		}
	case "reply":
		res.Type = "reply"
		if _, ok := seg.Data["message_id"]; ok && seg.Data["message_id"] != "" {
			res.Data["id"] = seg.Data["id"]
		}
		if _, ok := seg.Data["user_id"]; ok && seg.Data["user_id"] != "" {
			res.Data["qq"] = seg.Data["user_id"]
		}
	default:
		res.Type = seg.Type
		res.Data = seg.Data
	}
	return res
}

// Segment11To12 将v11的消息段转换为v12的消息段
func Segment11To12(seg protocol.Segment) protocol.Segment {
	var res = protocol.Segment{
		Data: make(map[string]interface{}),
	}
	switch seg.Type {
	case "at":
		res.Type = "mention"
		if _, ok := seg.Data["qq"]; ok && seg.Data["qq"] != "" {
			if seg.Data["qq"] == "all" {
				res.Type = "mention_all"
			} else {
				res.Data["user_id"] = seg.Data["qq"]
			}
		}
	case "record":
		res.Type = "voice"
		if _, ok := seg.Data["file"]; ok && seg.Data["file"] != "" {
			res.Data["file_id"] = seg.Data["file"]
		}
	case "video":
		res.Type = "video"
		if _, ok := seg.Data["file"]; ok && seg.Data["file"] != "" {
			res.Data["file_id"] = seg.Data["file"]
		}
	case "image":
		res.Type = "image"
		if _, ok := seg.Data["file"]; ok && seg.Data["file"] != "" {
			res.Data["file_id"] = seg.Data["file"]
		}
	case "location":
		res.Type = "location"
		if _, ok := seg.Data["lat"]; ok && seg.Data["lat"] != "" {
			res.Data["latitude"] = seg.Data["lat"]
		}
		if _, ok := seg.Data["lon"]; ok && seg.Data["lon"] != "" {
			res.Data["longitude"] = seg.Data["lon"]
		}
		if _, ok := seg.Data["title"]; ok && seg.Data["title"] != "" {
			res.Data["title"] = seg.Data["title"]
		}
		if _, ok := seg.Data["content"]; ok && seg.Data["content"] != "" {
			res.Data["content"] = seg.Data["content"]
		}
	case "reply":
		res.Type = "reply"
		if _, ok := seg.Data["id"]; ok && seg.Data["id"] != "" {
			res.Data["message_id"] = seg.Data["id"]
		}
		if _, ok := seg.Data["qq"]; ok && seg.Data["qq"] != "" {
			res.Data["user_id"] = seg.Data["qq"]
		}
	default:
		res.Type = seg.Type
		res.Data = seg.Data
	}
	return res
}

// CQCodeToSegment11 CQ码转11
func CQCodeToSegment11(s string) protocol.Segment {
	seg := protocol.Segment{
		Data: make(map[string]interface{}),
	}
	// 删除空格 前后括号
	s = deleteFromString(s, " ")
	s = deleteFromString(s, "[")
	s = deleteFromString(s, "]")
	// 以逗号分割
	str := strings.Split(s, ",")
	for _, v := range str {
		if strings.Index(v, "CQ:") == 0 && len(v) > 3 {
			seg.Type = v[3:]
		} else {
			if k := strings.Split(v, "="); len(k) == 2 {
				seg.Data[k[0]] = k[1]
			} else if len(k) == 1 {
				seg.Data[k[0]] = ""
			}
		}
	}
	return seg
}

// Segment11ToCQCode 11转CQ码
// TODO: 未完成
func Segment11ToCQCode(seg protocol.Segment) string {
	return ""
}

// CQCodeToSegment12 CQ码转12
func CQCodeToSegment12(s string) protocol.Segment {
	return Segment11To12(CQCodeToSegment11(s))
}

// Segment12ToCQCode 12转CQ码
// TODO: 未完成
func Segment12ToCQCode(seg protocol.Segment) string {
	return Segment11ToCQCode(Segment12To11(seg))
}

// decodeString 反转义
func decodeString(s string) string {
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, "&#91;", "[")
	s = strings.ReplaceAll(s, "&#93;", "]")
	s = strings.ReplaceAll(s, "&#44;", ",")
	return s
}

// Message11ToSegment12 将v11的消息转换为v12的消息段
func Message11ToSegment12(s string) []protocol.Segment {
	res := make([]protocol.Segment, 0)
	for len(s) != 0 {
		// 查找CQ码
		cqFrom := strings.Index(s, "[CQ:")
		cqTo := strings.Index(s, "]")
		if cqFrom == -1 || cqTo == -1 {
			// 找不到
			res = append(res, protocol.Segment{
				Type: "text",
				Data: map[string]interface{}{
					"text": decodeString(s),
				},
			})
			break
		} else {
			if cqTo < cqFrom {
				// 错误的cq码, 丢弃
				res = append(res, protocol.Segment{
					Type: "text",
					Data: map[string]interface{}{
						"text": decodeString(s[:cqTo+1]),
					},
				})
				s = s[cqTo+1:]
				continue
			}
			if cqFrom != 0 {
				// 前面的文字
				res = append(res, protocol.Segment{
					Type: "text",
					Data: map[string]interface{}{
						"text": decodeString(s[:cqFrom]),
					},
				})
			}
			res = append(res, CQCodeToSegment12(s[cqFrom:cqTo+1]))
			s = s[cqTo+1:]
		}
	}
	return res
}
