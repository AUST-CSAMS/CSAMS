package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// MajorArray 自定义类型用于将 []Major 转换为 JSON 存入数据库
type MajorArray []Major

// Scan 用于从数据库中将 JSON 数据解码为 MajorArray
func (m *MajorArray) Scan(value interface{}) error {
	// 检查 value 是否为 nil
	if value == nil {
		*m = nil
		return nil
	}

	// 将 value 转换为 []byte
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("unsupported data type for MajorArray")
	}

	// 解码 JSON 数据到 CustomIntSlice
	return json.Unmarshal(bytes, m)
}

// Value 用于将 MajorArray 转换为数据库中的 JSON 数据
func (m MajorArray) Value() (driver.Value, error) {
	// 将 CustomIntSlice 转换为 JSON 数据
	return json.Marshal(m)
}
