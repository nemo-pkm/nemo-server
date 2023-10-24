package scripts

import (
	"bufio"
	"fmt"
	"os"
)

func DealArticleContent(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	text := []string{}
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	var str string
	for _, v := range text {
		if len(v) == 0 {
			str += "\n"
			continue
		}
		str += v + "\n"
	}
	// 检查是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件错误:", err)
	}
	return str
}

func DealPostContent(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	text := []string{}
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	var str string
	for _, v := range text {
		if len(v) == 0 {
			str += "\n"
			continue
		}
		str += v + "\n"
	}
	// 检查是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件错误:", err)
	}
	return str
}
