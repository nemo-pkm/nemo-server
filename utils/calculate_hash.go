package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

/*
	// 示例文本
	text := "Hello, World!"

	// 计算文本的哈希值
	hashValue := utils.CalculateTextHash(text)

	// 输出哈希值
	fmt.Println("文本的SHA-256哈希值为:", hashValue)
*/

// 计算文本的哈希值
func CalculateTextHash(text string) string {
	// 将文本转换为字节序列
	encodedText := []byte(text)

	// 创建SHA-256哈希对象
	hash := sha256.New()

	// 计算哈希值
	hash.Write(encodedText)
	hashValue := hash.Sum(nil)

	// 将哈希值转换为十六进制表示
	hashString := hex.EncodeToString(hashValue)

	return hashString
}
