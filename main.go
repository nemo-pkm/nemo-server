package main

import (
	"fmt"
	"nemo/scripts"
	"time"
)

//	type notes struct {
//		Content    string `json:content`
//		Visibility string `json:visibility`
//		CreatedTs  int64  `json:createdTs`
//		UpdateTs   int64  `json:updateTs`
//		DisplayTs  int64  `json:displayTs`
//	}
// type SingleMemo struct {
// 	Content    string `json:"content"`
// 	Visibility string `json:"visibility"`
// 	CreatedTs  int64  `json:"createdTs"`
// 	UpdateTs   int64  `json:"updateTs"`
// 	DisplayTs  int64  `json:"displayTs"`
// }

func main() {
	// note := scripts.WechatReadingNotes{
	// 	NotesFilePath: "C:\\file\\git\\nemo\\wechat.txt",
	// 	MemoTime:      "2021-4-14 16:00:00",
	// 	Tags:          []string{"灵山", "微信读书"},
	// 	Url:           "http://103.82.52.107:5230/api/v1/memo",
	// 	Visibility:    "PRIVATE",
	// 	Token:         "eyJhbGciOiJIUzI1NiIsImtpZCI6InYxIiwidHlwIjoiSldUIn0.eyJuYW1lIjoibmVvIiwiaXNzIjoibWVtb3MiLCJzdWIiOiIxIiwiYXVkIjpbInVzZXIuYWNjZXNzLXRva2VuIl0sImV4cCI6MTY5ODY2ODM5NSwiaWF0IjoxNjk4MDYzNTk1fQ.Bb5RZoHNbvgLp5olg_5sQ0Ic2XN51fB4PP4fC6GKaNg",
	// }
	// note.ImportWeChatReadingNotes()
	timestr := "2023-10-24 7:00:00"
	//neo := scripts.DealPostContent("C:\\file\\git\\nemo\\content.txt") + "\n#memo"
	neo := scripts.DealArticleContent("C:\\file\\git\\nemo\\content.txt") + "\n#memo"
	memo := scripts.SingleMemo{
		Visibility: "PRIVATE",
	}
	memo.Content = neo
	memo.CreatedTs = timeToUnixTime(timestr)
	memo.DisplayTs = timeToUnixTime(timestr)
	url := "http://103.82.52.107:5230/api/v1/memo"
	token := "eyJhbGciOiJIUzI1NiIsImtpZCI6InYxIiwidHlwIjoiSldUIn0.eyJuYW1lIjoibmVvIiwiaXNzIjoibWVtb3MiLCJzdWIiOiIxIiwiYXVkIjpbInVzZXIuYWNjZXNzLXRva2VuIl0sImV4cCI6MTY5ODY2ODM1NSwiaWF0IjoxNjk4MDYzNTU1fQ.ujK4qICgYDQV4m-C0GJmuD73nd7VWGc7KFkwGtoZgeA"
	memo.ImportSingleMemo(url, token)

}
func timeToUnixTime(timeString string) int64 {
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
