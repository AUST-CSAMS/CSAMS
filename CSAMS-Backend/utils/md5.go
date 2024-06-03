package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 md5
func Md5(src []byte) string {
	//创建一个md5实例
	m := md5.New()
	//写入数据
	m.Write(src)
	//m.Sum(nil)返回结果，hex.EncodeToString转换16进制
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
