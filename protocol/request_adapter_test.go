package protocol

import (
	"github.com/FishZe/go-libonebot/util"
	"reflect"
	"testing"
)

type requestAdapterArgs struct {
	e any
	r RawRequestType
}
type requestAdapterWant struct {
	wantErr error
}

type requestAdapterTest struct {
	name      string
	args      requestAdapterArgs
	want      requestAdapterWant
	checkFunc func(t *testing.T, e any, r RawRequestType) bool
}

func requestTestCheck(t *testing.T, e any, r RawRequestType) bool {
	tt := reflect.TypeOf(e).Elem()
	vv := reflect.ValueOf(e).Elem()
	for i := 0; i < tt.NumField(); i++ {
		if tt.Field(i).Type == RequestType {
			if vv.Field(i).Interface().(*Request).Action != r.Action {
				t.Errorf("RequestAdapter() = %v, want %v", vv.Field(i).Interface().(*Request).Action, r.Action)
				return false
			}
			if vv.Field(i).Interface().(*Request).Echo != r.Echo {
				t.Errorf("RequestAdapter() = %v, want %v", vv.Field(i).Interface().(*Request).Echo, r.Echo)
				return false
			}
			if vv.Field(i).Interface().(*Request).Self != r.Self {
				t.Errorf("RequestAdapter() = %v, want %v", vv.Field(i).Interface().(*Request).Self, r.Self)
				return false
			}
			if vv.Field(i).Interface().(*Request).ConnectionUUID != r.ConnectionUUID {
				t.Errorf("RequestAdapter() = %v, want %v", vv.Field(i).Interface().(*Request).ConnectionUUID, r.ConnectionUUID)
				return false
			}
		}
	}
	return false
}

func getRequestAdapterTests() []requestAdapterTest {
	res := make([]requestAdapterTest, 0)
	res = append(res, requestAdapterTest{
		name: "Request test error 1",
		args: requestAdapterArgs{
			e: nil,
			r: RawRequestType{},
		},
		want: requestAdapterWant{wantErr: ErrorInvalidRequest},
	})
	res = append(res, requestAdapterTest{
		name: "Request test error 2",
		args: requestAdapterArgs{
			e: &struct{}{},
			r: RawRequestType{},
		},
		want: requestAdapterWant{wantErr: ErrorInvalidRequest},
	})
	res = append(res, requestAdapterTest{
		name: "Request test error 3",
		args: requestAdapterArgs{
			e: 123,
			r: RawRequestType{},
		},
		want: requestAdapterWant{wantErr: ErrorInvalidRequest},
	})
	res = append(res, requestAdapterTest{
		name: "Request  test error 4",
		args: requestAdapterArgs{
			e: &struct {
				Msg string `json:"msg"`
			}{
				Msg: "test",
			},
			r: RawRequestType{},
		},
		want: requestAdapterWant{wantErr: ErrorInvalidRequest},
	})
	r := RequestGetUserInfo{}
	res = append(res, requestAdapterTest{
		name: "Request test error 5",
		args: requestAdapterArgs{
			e: &r,
			r: RawRequestType{
				&Request{},
				make(map[string]any),
			},
		},
		want: requestAdapterWant{wantErr: ErrorRequestIsNil},
	})
	r1 := r.New()
	res = append(res, requestAdapterTest{
		name: "Request test error 6",
		args: requestAdapterArgs{
			e: r1,
			r: RawRequestType{
				Request: &Request{
					Action: "",
				},
			},
		},
		want:      requestAdapterWant{wantErr: ErrorActionEmpty},
		checkFunc: requestTestCheck,
	})
	r2 := r.New()
	res = append(res, requestAdapterTest{
		name: "Request test error 7",
		args: requestAdapterArgs{
			e: r2,
			r: RawRequestType{
				Request: &Request{
					Action: EventTypeMeta,
				},
			},
		},
		want:      requestAdapterWant{wantErr: ErrorRequestNotMatch},
		checkFunc: requestTestCheck,
	})
	r3 := r.New()
	res = append(res, requestAdapterTest{
		name: "Request test normal 1",
		args: requestAdapterArgs{
			e: r3,
			r: RawRequestType{
				Request: &Request{
					Action:         "get_user_info",
					Echo:           "test",
					Self:           Self{"test", "test"},
					ConnectionUUID: util.GetUUID(),
					requestID:      util.GetUUID(),
				},
			},
		},
		want:      requestAdapterWant{wantErr: nil},
		checkFunc: requestTestCheck,
	})
	rr := RequestUploadFileFragmented{}
	r4 := rr.New()
	param4 := make(map[string]any)
	param4["stage"] = "transfer"
	param4["name"] = "helllo"
	param4["total_size"] = 123
	param4["sha256"] = util.GetUUID()
	param4["offset"] = 10
	param4["file_id"] = util.GetUUID()
	param4["data"] = []byte(util.GetUUID())
	for i := 0; i <= 100; i++ {
		param4["data"] = append(param4["data"].([]byte), []byte(util.GetUUID())...)
	}
	res = append(res, requestAdapterTest{
		name: "Request test normal 2",
		args: requestAdapterArgs{
			e: r4,
			r: RawRequestType{
				Request: &Request{
					Action:         "upload_file_fragmented",
					Echo:           "test",
					Self:           Self{"test", "test"},
					ConnectionUUID: util.GetUUID(),
					requestID:      util.GetUUID(),
				},
				Param: param4,
			},
		},
		want:      requestAdapterWant{wantErr: nil},
		checkFunc: requestTestCheck,
	})
	return res
}

func BenchmarkRequestAdapter(b *testing.B) {
	util.Logger.SetLogLevel(util.LogLevelInfo)
	tests := getRequestAdapterTests()
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = RequestAdapter(tt.args.e, tt.args.r)
			}
		})
	}
	util.Logger.SetLogLevel(util.LogLevelWarning)
}

func TestRequestAdapter(t *testing.T) {
	tests := getRequestAdapterTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RequestAdapter(tt.args.e, tt.args.r); err != nil {
				if err != tt.want.wantErr {
					t.Errorf("RequestAdapter() error = %v, wantErr %v", err, tt.want.wantErr)
				}
			} else if tt.want.wantErr != nil {
				t.Errorf("RequestAdapter() error = %v, wantErr %v", err, tt.want.wantErr)
			} else if tt.checkFunc != nil {
				tt.checkFunc(t, tt.args.e, tt.args.r)
			}
		})
	}
}
