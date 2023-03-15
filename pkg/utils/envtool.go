package utils

import (
	"os"
)

// GetEnv 获取环境变量的值，如果环境变量不存在则返回默认值
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}

// SetEnv 设置环境变量的值
func SetEnv(key, value string) {
	os.Setenv(key, value)
}
