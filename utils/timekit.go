package utils

import (
	"fmt"
	"time"
)

// 最终便签显式哪个时区在前端页面上选择

func TimeToUnixTime(timeString string) int64 {
	// 时间格式化模板
	timeLayout := "2006-01-02 15:04:05"
	// 解析时间字符串
	parsedTime, err := time.Parse(timeLayout, timeString)
	if err != nil {
		fmt.Println("时间解析错误:", err)
		return 0
	}
	return parsedTime.Unix()
}
