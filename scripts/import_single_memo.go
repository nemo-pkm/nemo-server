package scripts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type SingleMemo struct {
	Content    string `json:"content"`
	Visibility string `json:"visibility"`
	CreatedTs  int64  `json:"createdTs"`
	UpdateTs   int64  `json:"updateTs"`
	DisplayTs  int64  `json:"displayTs"`
}

func (memo SingleMemo) ImportSingleMemo(url string, token string) {
	jsondata, err := json.Marshal(memo)
	if err != nil {
		fmt.Println("JSON编码失败:", err)
		return
	}
	fmt.Println(string(jsondata))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsondata))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}
	req.Header.Set("Content-Length", strconv.Itoa(len(jsondata)))
	req.Header.Set("Authorization", "Bearer "+token)
	// 发送请求
	client := &http.Client{}
	for attempt := 1; attempt <= 6; attempt++ {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("发送请求失败:", err)
			continue
		}
		// 读取响应结果
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取响应失败:", err)
			resp.Body.Close()
			continue
		}

		// 打印响应结果
		fmt.Println("响应结果:", string(body))
		resp.Body.Close()
		break // 请求成功，跳出重试循环
	}

}
