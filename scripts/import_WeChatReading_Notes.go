package scripts

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"nemo/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

/*
filepath 文件路径

memoCreateTime memo 时间
时间样例
timeLayout := "2006-01-02 15:04:05"

tags memo 的标签
*/

type WechatReadingNotes struct {
	NotesFilePath string
	MemoTime      string
	Visibility    string
	Tags          []string
	Url           string
	Token         string
}

func NewWechatReadingNotes(notesFilePath string, memoTime string, visibility string, tags []string, url string, token string) *WechatReadingNotes {
	return &WechatReadingNotes{NotesFilePath: notesFilePath, MemoTime: memoTime, Visibility: visibility, Tags: tags, Url: url, Token: token}
}

// 每一条 memo 的结构
type ImportMemo struct {
	Content    string `json:"content"`
	Visibility string `json:"visibility"`
	CreatedTs  int64  `json:"createdTs"`
	UpdateTs   int64  `json:"updateTs"`
	DisplayTs  int64  `json:"displayTs"`
}

func NewImportMemo(content string, visibility string, Time int64) *ImportMemo {
	return &ImportMemo{Content: content, Visibility: visibility, CreatedTs: Time, UpdateTs: Time, DisplayTs: Time}
}

type ImportTag struct {
	Name string `json:"name"`
}

func NewImportTag(name string) *ImportTag {
	return &ImportTag{Name: name}
}

// 导入函数
func (notes WechatReadingNotes) ImportWeChatReadingNotes() {
	notes.MergeTags()
	file, err := os.Open(notes.NotesFilePath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	text := []string{}
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// 检查是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件错误:", err)
	}
	var memolist []ImportMemo
	for _, v := range text {
		v = strings.ReplaceAll(v, "\n", "")
		if len(v) >= 1 {
			if string(v[0]) == ">" {
				memoContent := v[3:] + "\n\n" + genrateTags(notes.Tags)
				memo := NewImportMemo(memoContent, notes.Visibility, timeToUnixTime(notes.MemoTime))
				memolist = append(memolist, *memo)
			}
		}
	}
	strlist := []string{}
	for _, memo := range memolist {
		jsondata, err := json.Marshal(memo)
		if err != nil {
			fmt.Println("JSON编码失败:", err)
			return
		}
		str := string(jsondata)
		strlist = append(strlist, str)
	}
	strjson := "[" + strings.Join(strlist, ",") + "]"
	WriteJSONToFile(strjson, "C:\\file\\git\\nemo\\wechat.json")
	fmt.Println("==================sendOS======================")
	sendPOST(memolist, notes.Url, notes.Token)
}

func genrateTags(strSlice []string) string {
	var str string
	for _, s := range strSlice {
		str += "#" + string(s) + " "
	}
	return str[:len(str)-1]
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

func sendPOST(memolist []ImportMemo, url string, token string) {
	for i, memo := range memolist {
		fmt.Println(i)
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
}

func (notes WechatReadingNotes) MergeTags() {
	// 创建GET请求
	req, err := http.NewRequest("GET", "http://103.82.52.107:5230/api/v1/tag", nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+notes.Token)
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应结果
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}
	fmt.Println(string(body))

	// 解析JSON为切片
	result := gjson.ParseBytes(body)
	listExitTags := make([]string, 0)

	result.ForEach(func(key, value gjson.Result) bool {
		listExitTags = append(listExitTags, value.String())
		return true
	})

	fmt.Println(listExitTags)

	Tags := utils.RemoveDuplicates(notes.Tags, listExitTags)
	if len(Tags) == 0 {
		return
	}
	fmt.Println(Tags)
	for i, tag := range Tags {
		println(i)
		addtag := NewImportTag(tag)
		jsondata, err := json.Marshal(addtag)
		if err != nil {
			fmt.Println("JSON编码失败:", err)
			return
		}

		fmt.Println(string(jsondata))

		req, err := http.NewRequest("POST", "http://103.82.52.107:5230/api/v1/tag", bytes.NewBuffer(jsondata))
		if err != nil {
			fmt.Println("创建请求失败:", err)
			return
		}
		req.Header.Set("Authorization", "Bearer "+notes.Token)

		// 发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("发送请求失败:", err)
			return
		}
		// 读取响应结果
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取响应失败:", err)
			return
		}

		// 打印响应结果
		fmt.Println("响应结果:", string(body))
		resp.Body.Close()
		time.Sleep(3000 * time.Millisecond)

	}

}

func WriteJSONToFile(jsonStr string, filePath string) error {
	return os.WriteFile(filePath, []byte(jsonStr), 0644)
}
