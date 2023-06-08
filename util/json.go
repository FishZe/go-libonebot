package util

import "github.com/bytedance/sonic"

// JsonCoder 一个Json解析器
type jsonCoder interface {
	// Unmarshal 解码
	Unmarshal(data []byte, v interface{}) error
	// Marshal 编码
	Marshal(v interface{}) ([]byte, error)
}

// Json 使用的解析器
// 使用字节的sonic库 由于不支持arm, 所以提供了修改的方法
var Json jsonCoder

// DefaultJson 默认解析器
type DefaultJson struct {
}

// Unmarshal 默认解析器解码
func (d *DefaultJson) Unmarshal(data []byte, v interface{}) error {
	return sonic.Unmarshal(data, v)
}

// Marshal 默认解析器编码
func (d *DefaultJson) Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

// SetJsonCoder 修改Json解析器
func SetJsonCoder(j jsonCoder) {
	Json = j
}

// init 初始化 设置为默认解析器
func init() {
	// 默认使用sonic
	SetJsonCoder(new(DefaultJson))
	// 使用默认日志
	SetLogger(new(Log))
	// 等级
	Logger.SetLogLevel(LogLevelDebug)
}
