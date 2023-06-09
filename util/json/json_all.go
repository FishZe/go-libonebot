//go:build !amd64

package json

// DefaultJson 默认解析器
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
