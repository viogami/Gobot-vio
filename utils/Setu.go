package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type SetuRequest struct {
	R18        int      `json:"r18"`
	Num        int      `json:"num"`
	Uid        []int    `json:"uid"`
	Tag        []string `json:"tag"`
	Size       []string `json:"size"`
	Proxy      string   `json:"proxy"`
	DateAfter  int64    `json:"dateAfter"`
	DateBefore int64    `json:"dateBefore"`
	DSC        bool     `json:"dsc"`
	ExcludeAI  bool     `json:"excludeAI"`
}

type SetuResponse struct {
	Error string     `json:"error"`
	Data  []SetuData `json:"data"`
}
type SetuData struct {
	Pid        int      `json:"pid"`
	P          int      `json:"p"`
	Uid        int      `json:"uid"`
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	R18        bool     `json:"r18"`
	Width      int      `json:"width"`
	Height     int      `json:"height"`
	Tags       []string `json:"tags"`
	Ext        string   `json:"ext"`
	AiType     int      `json:"aiType"`
	UploadDate int64    `json:"uploadDate"`
	Urls       Urls     `json:"urls"`
}

type Urls struct {
	Original string `json:"original"`
	Regular  string `json:"regular"`
	Small    string `json:"small"`
}

func Get_setu(tags []string, r18 int, num int) SetuResponse {
	// 示例：构造一个 SetuRequest
	requestData := SetuRequest{
		R18:       r18,
		Num:       num,
		Tag:       tags,
		Size:      []string{"small", "regular"},
		Proxy:     "i.pixiv.re",
		DSC:       false,
		ExcludeAI: false,
	}
	// 示例：调用 Setu API
	// 将请求参数转换为 JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}

	// 发送 POST 请求
	url := "https://api.lolicon.app/setu/v2"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body := new(bytes.Buffer)
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	// 解析 JSON 响应
	var setuResponse SetuResponse
	err = json.Unmarshal(body.Bytes(), &setuResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	return setuResponse
}

// 判断是否发送涩图
func Get_tags(input string) []string {
	// 检查是否以 "/涩图 " 开头
	if strings.HasPrefix(input, "/涩图 ") {
		// 获取 "/涩图 " 后面的部分
		tagsPart := strings.TrimPrefix(input, "/涩图 ")
		// 使用正则表达式匹配中英文逗号
		tags := strings.Split(tagsPart, "，")
		return tags
	}
	return []string{}
}
