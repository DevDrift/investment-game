package utils

import (
	"os"
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
