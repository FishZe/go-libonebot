package protocol

import (
	"github.com/FishZe/go-libonebot/util"
	"testing"
)

// 测试文件不写注释了，累了

type eventCheckArgs struct {
	s Self
	e any
}
type eventCheckWant struct {
	want    string
	want1   string
	wantErr error
}

type eventCheckTest struct {
	name      string
	args      eventCheckArgs
	want      eventCheckWant
	checkFunc func(t *testing.T, string2 string) bool
}

func getEventCheckTests() []eventCheckTest {
	res := make([]eventCheckTest, 0)
	res = append(res, eventCheckTest{
		name: "Event test error 1",
		args: eventCheckArgs{
			s: Self{"platform", "123456"},
			e: nil,
		},
		want: eventCheckWant{"", "", ErrorInvalidEvent},
	})
	res = append(res, eventCheckTest{
		name: "Event test error 2",
		args: eventCheckArgs{
			s: Self{"platform", "123456"},
			e: struct{}{},
		},
		want: eventCheckWant{"", "", ErrorInvalidEvent},
	})
	res = append(res, eventCheckTest{
		name: "Event test error 3",
		args: eventCheckArgs{
			s: Self{"platform", "123456"},
			e: &struct{}{},
		},
		want: eventCheckWant{"", "", ErrorInvalidEvent},
	})
	res = append(res, eventCheckTest{
		name: "Event test error 4",
		args: eventCheckArgs{
			s: Self{"platform", "123456"},
			e: &struct {
				Msg string `json:"msg"`
			}{
				Msg: "test",
			},
		},
		want: eventCheckWant{"", "", ErrorInvalidEvent},
	})
	res = append(res, eventCheckTest{
		name: "Event test error 5",
		args: eventCheckArgs{
			s: Self{"platform", "123456"},
			e: &struct {
				Event *Event `json:"event"`
				Msg   string `json:"msg"`
			}{
				Event: nil,
				Msg:   "test",
			},
		},
		want: eventCheckWant{"", "", ErrorInvalidEvent},
	})
	res = append(res, eventCheckTest{
		name: "Event test error 6",
		args: eventCheckArgs{
			s: Self{"platform", "123456"},
			e: &struct {
				Event *Event `json:"event"`
				Msg   string `json:"msg"`
			}{
				Event: &Event{},
				Msg:   "test",
			},
		},
		want: eventCheckWant{"", "", ErrorInValidEventType},
	})
	res = append(res, eventCheckTest{
		name: "Event test error 7",
		args: eventCheckArgs{
			s: Self{"platform", "123456"},
			e: &struct {
				Event *Event `json:"event"`
				Msg   string `json:"msg"`
			}{
				Event: &Event{
					Type: "testType",
				},
				Msg: "test",
			},
		},
		want: eventCheckWant{"", "", ErrorInValidEventType},
	})
	nowEvent1 := Event{
		Type: EventTypeMeta,
	}
	res = append(res, eventCheckTest{
		name: "Event test normal 1",
		args: eventCheckArgs{
			s: Self{"platform", "123456"},
			e: &struct {
				Event *Event `json:"event"`
				Msg   string `json:"msg"`
			}{
				Event: &nowEvent1,
				Msg:   "test",
			},
		},
		want: eventCheckWant{wantErr: nil},
		checkFunc: func(t *testing.T, _ string) bool {
			if nowEvent1.ID == "" {
				t.Errorf("EventCheck() ID not set")
				return false
			}
			if nowEvent1.Time == 0 {
				t.Errorf("EventCheck() Time not set")
				return false
			}
			return true
		},
	})
	nowEvent2 := Event{
		Type: EventTypeMeta,
		Time: util.GetTimeStampFloat64(),
		ID:   "123456",
	}
	nowEvent2Time := nowEvent2.Time
	res = append(res, eventCheckTest{
		name: "Event ID & Time been covered",
		args: eventCheckArgs{
			s: Self{"platform", "123456"},
			e: &struct {
				Event *Event `json:"event"`
				Msg   string `json:"msg"`
			}{
				Event: &nowEvent2,
				Msg:   "test",
			},
		},
		want: eventCheckWant{wantErr: nil},
		checkFunc: func(t *testing.T, _ string) bool {
			if nowEvent2.ID != "123456" {
				t.Errorf("EventCheck() ID was covered")
				return false
			}
			if nowEvent2.Time != nowEvent2Time {
				t.Errorf("EventCheck() Time was covered")
				return false
			}
			return true
		},
	})
	nowEvent3 := Event{
		Type: EventTypeMeta,
	}
	nowEvent3Self := Self{"platform", "123456"}
	res = append(res, eventCheckTest{
		name: "Event self auto set",
		args: eventCheckArgs{
			s: nowEvent3Self,
			e: &struct {
				Event *Event `json:"event"`
				Msg   string `json:"msg"`
			}{
				Event: &nowEvent3,
				Msg:   "test",
			},
		},
		want: eventCheckWant{wantErr: nil},
		checkFunc: func(t *testing.T, _ string) bool {
			if nowEvent3.Self != nowEvent3Self {
				t.Errorf("EventCheck() Self not set")
				return false
			}
			return true
		},
	})
	nowEvent4 := Event{
		Type:       EventTypeMeta,
		DetailType: "testDetailType",
	}
	res = append(res, eventCheckTest{
		name: "Event self auto set",
		args: eventCheckArgs{
			s: Self{"platform", "123456"},
			e: &struct {
				Event *Event `json:"event"`
				Msg   string `json:"msg"`
			}{
				Event: &nowEvent4,
				Msg:   "test",
			},
		},
		want: eventCheckWant{wantErr: nil},
		checkFunc: func(t *testing.T, s string) bool {
			if s == EventTypeMeta+"/"+"testDetailType" {
				return false
			}
			return true
		},
	})
	return res
}

func BenchmarkEventCheck(b *testing.B) {
	// 防止IO时间影响
	util.Logger.SetLogLevel(util.LogLevelInfo)
	tests := getEventCheckTests()
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = EventCheck(tt.args.s, tt.args.e)
			}
		})
	}
	util.Logger.SetLogLevel(util.LogLevelWarning)
}

func TestEventCheck(t *testing.T) {
	util.Logger.SetLogLevel(util.LogLevelInfo)
	tests := getEventCheckTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := EventCheck(tt.args.s, tt.args.e)
			if err != tt.want.wantErr {
				t.Errorf("EventCheck() error = %v, wantErr %v", err, tt.want.wantErr)
				return
			} else {
				if tt.checkFunc != nil {
					x := tt.checkFunc(t, got)
					if !x {
						return
					}
				}
			}
		})
	}
}
