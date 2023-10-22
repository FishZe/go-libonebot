package util

import "github.com/goccy/go-json"

// JsonCoder 一个Json解析器
type jsonCoder interface {
	// Unmarshal 解码
	Unmarshal(data []byte, v interface{}) error
	// Marshal 编码
	Marshal(v interface{}) ([]byte, error)
}

type DefaultJson struct {
}

// Unmarshal 默认解析器解码
func (d *DefaultJson) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Marshal 默认解析器编码
func (d *DefaultJson) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Json 使用的解析器
// 使用字节的sonic库 由于只支持amd64, 所以提供了修改的方法
var Json jsonCoder

// SetJsonCoder 修改Json解析器
func SetJsonCoder(j jsonCoder) {
	Json = j
}

// init 初始化 设置为默认解析器
func init() {
	// 默认使用sonic
	SetJsonCoder(new(DefaultJson))
}
