package utils

import (
	"os"

	"github.com/bytedance/sonic"
)

// GetEnv - возвращает значение переменной окружения для указанного ключа.
// Если переменная окружения не установлена, возвращается значение по умолчанию (fallback).
func GetEnv(key, fallback string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// ToJsonBytes конвертация в массив байт
func ToJsonBytes(input interface{}) []byte {
	marshal, err := sonic.Marshal(input)
	if err != nil {
		return []byte(``)
	}
	return marshal
}

// ToBytesJson конвертация из массива байт
func ToBytesJson(input []byte) (interface{}, error) {
	var res interface{}
	err := sonic.Unmarshal(input, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
