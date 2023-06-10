package protocol

import (
	"github.com/FishZe/go-libonebot/util"
	"reflect"
	"testing"
)

type responseAdapterArgs struct {
	r *Request
	e any
}
type responseAdapterWant struct {
	wantErr error
}

type responseAdapterTest struct {
	name      string
	args      responseAdapterArgs
	want      responseAdapterWant
	checkFunc func(t *testing.T, r *Request, e any) bool
}

func getResponseAdapterTests() []responseAdapterTest {
	res := make([]responseAdapterTest, 0)
	res = append(res, responseAdapterTest{
		name: "Response test error 1",
		args: responseAdapterArgs{
			r: nil,
			e: nil,
		},
		want: responseAdapterWant{ErrorInvalidResponse},
	})
	res = append(res, responseAdapterTest{
		name: "Response test error 2",
		args: responseAdapterArgs{
			r: &Request{},
			e: 123,
		},
		want: responseAdapterWant{ErrorInvalidResponse},
	})
	res = append(res, responseAdapterTest{
		name: "Response test error 3",
		args: responseAdapterArgs{
			r: &Request{},
			e: &struct{}{},
		},
		want: responseAdapterWant{ErrorInvalidResponse},
	})
	res = append(res, responseAdapterTest{
		name: "Response test error 4",
		args: responseAdapterArgs{
			r: &Request{},
			e: &struct {
				Msg string `json:"msg"`
			}{
				Msg: "test",
			},
		},
		want: responseAdapterWant{ErrorInvalidResponse},
	})
	res = append(res, responseAdapterTest{
		name: "Response test error 5",
		args: responseAdapterArgs{
			r: &Request{},
			e: &struct {
				Response
			}{
				Response{},
			},
		},
		want: responseAdapterWant{ErrorInvalidResponse},
	})
	res = append(res, responseAdapterTest{
		name: "Response test error 6",
		args: responseAdapterArgs{
			r: &Request{},
			e: &struct {
				*Response
			}{
				&Response{},
			},
		},
		want: responseAdapterWant{nil},
	})
	res = append(res, responseAdapterTest{
		name: "Response test error 7",
		args: responseAdapterArgs{
			r: &Request{},
			e: &struct {
				*Response
			}{
				&Response{
					Retcode: -1,
				},
			},
		},
		want: responseAdapterWant{ErrorInvalidResponseRetCode},
	})
	rr := NewEmptyResponse(0)
	req := &Request{
		requestID: util.GetUUID(),
		Echo:      util.GetUUID(),
	}
	res = append(res, responseAdapterTest{
		name: "Response test normal 1",
		args: responseAdapterArgs{
			r: req,
			e: rr,
		},
		want:      responseAdapterWant{nil},
		checkFunc: responseCheckTest,
	})
	rr2 := NewEmptyResponse(10001)
	req2 := &Request{
		requestID: util.GetUUID(),
		Echo:      util.GetUUID(),
	}
	res = append(res, responseAdapterTest{
		name: "Response test normal 2",
		args: responseAdapterArgs{
			r: req2,
			e: rr2,
		},
		want:      responseAdapterWant{nil},
		checkFunc: responseCheckTest,
	})
	rr3 := NewResponseGetFile(0)
	rr3.Url = util.GetUUID()
	rr3.Data = []byte(util.GetUUID())
	for i := 0; i <= 100; i++ {
		rr3.Data = append(rr3.Data, []byte(util.GetUUID())...)
	}
	req3 := &Request{
		requestID: util.GetUUID(),
		Echo:      util.GetUUID(),
	}
	res = append(res, responseAdapterTest{
		name: "Response test normal 3",
		args: responseAdapterArgs{
			r: req3,
			e: rr3,
		},
		want:      responseAdapterWant{nil},
		checkFunc: responseCheckTest,
	})
	return res
}

func responseCheckTest(t *testing.T, r *Request, e any) bool {
	tt := reflect.TypeOf(e).Elem()
	vv := reflect.ValueOf(e).Elem()
	for i := 0; i < tt.NumField(); i++ {
		if tt.Field(i).Type == ResponseType {
			if vv.Field(i).Interface().(*Response).Retcode == 0 && vv.Field(i).Interface().(*Response).Status != StatusOk {
				t.Errorf("RequestAdapter() = %v, want %v", vv.Field(i).Interface().(*Request).Action, r.Action)
				return false
			}
			if vv.Field(i).Interface().(*Response).Retcode != 0 && vv.Field(i).Interface().(*Response).Status != StatusFailed {
				t.Errorf("RequestAdapter() = %v, want %v", vv.Field(i).Interface().(*Request).Action, r.Action)
				return false
			}
			if vv.Field(i).Interface().(*Response).Echo != r.Echo {
				t.Errorf("RequestAdapter() = %v, want %v", vv.Field(i).Interface().(*Request).Echo, r.Echo)
				return false
			}
			if vv.Field(i).Interface().(*Response).requestID != r.GetID() {
				t.Errorf("RequestAdapter() = %v, want %v", vv.Field(i).Interface().(*Request).GetID(), r.GetID())
				return false
			}
		}
	}
	return false
}

func BenchmarkResponseCheck(b *testing.B) {
	util.Logger.SetLogLevel(util.LogLevelInfo)
	tests := getResponseAdapterTests()
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ResponseCheck(tt.args.r, tt.args.e)
			}
		})
	}
	util.Logger.SetLogLevel(util.LogLevelWarning)
}

func TestResponseCheck(t *testing.T) {
	util.Logger.SetLogLevel(util.LogLevelInfo)
	tests := getResponseAdapterTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ResponseCheck(tt.args.r, tt.args.e); err != nil {
				if err != tt.want.wantErr {
					t.Errorf("ResponseCheck() error = %v, wantErr %v", err, tt.want.wantErr)
				}
			} else if tt.want.wantErr != nil {
				t.Errorf("ResponseCheck() error = %v, wantErr %v", err, tt.want.wantErr)
			} else if tt.checkFunc != nil {

			}
		})
	}
}
