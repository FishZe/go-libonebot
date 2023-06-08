package v11_adapter

import (
	"github.com/FishZe/go-libonebot/protocol"
)

func Segment12To11(seg protocol.Segment) protocol.Segment {
	return protocol.Segment{}
}

func Segment11To12(seg protocol.Segment) protocol.Segment {
	return protocol.Segment{}
}

func CQCodeToSegment11(s string) protocol.Segment {
	return protocol.Segment{}
}

func Segment11ToCQCode(seg protocol.Segment) string {
	return ""
}

func CQCodeToSegment12(s string) protocol.Segment {
	return Segment11To12(CQCodeToSegment11(s))
}

func Segment12ToCQCode(seg protocol.Segment) string {
	return Segment11ToCQCode(Segment12To11(seg))
}
