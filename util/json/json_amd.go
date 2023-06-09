//go:build amd64

package json

import "github.com/bytedance/sonic"

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
