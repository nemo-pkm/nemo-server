package main

import "nemo/scripts"

type notes struct {
	Content    string `json:content`
	Visibility string `json:visibility`
	CreatedTs  int64  `json:createdTs`
	UpdateTs   int64  `json:updateTs`
	DisplayTs  int64  `json:displayTs`
}

func main() {
	note := scripts.WechatReadingNotes{
		NotesFilePath: "C:\\file\\git\\nemo\\wechat.txt",
		MemoTime:      "2021-4-14 16:00:00",
		Tags:          []string{"灵山", "微信读书"},
		Url:           "http://103.82.52.107:5230/api/v1/memo",
		Visibility:    "PRIVATE",
		Token:         "eyJhbGciOiJIUzI1NiIsImtpZCI6InYxIiwidHlwIjoiSldUIn0.eyJuYW1lIjoibmVvIiwiaXNzIjoibWVtb3MiLCJzdWIiOiIxIiwiYXVkIjpbInVzZXIuYWNjZXNzLXRva2VuIl0sImV4cCI6MTY5ODUzMDk5NSwiaWF0IjoxNjk3OTI2MTk1fQ.s0MFNSSEBa7Cf1V_oKYWgCytkizjxr39WdrYB-XiCaE",
	}
	note.ImportWeChatReadingNotes()
}
