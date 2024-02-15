package uuid

import "github.com/google/uuid"

// GenerateUUID UUID生成器
// 传入:无
// 传出:UUID
func GenerateUUID() string {
	return uuid.New().String()
}
