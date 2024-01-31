package gocq

import (
	"fmt"
	"regexp"
	"strings"
)

type CQmsg struct {
	CQcodes []CQCode
	Text    string
}

type CQCode struct {
	Type   string
	Params map[string]interface{}
}

func ParseCQmsg(input string) CQmsg {
	re := regexp.MustCompile(`\[CQ:([^,]+)(?:,([^=]+)=([^,]+))?\]`)

	matches := re.FindAllStringSubmatch(input, -1)

	result := CQmsg{}

	// 初始化 Text 为原始 input
	result.Text = input

	// 处理每个CQ码句段
	for _, match := range matches {
		cqCode := CQCode{
			Type:   match[1],
			Params: make(map[string]interface{}),
		}

		if match[2] != "" && match[3] != "" {
			// 如果有参数，将参数添加到 map 中
			cqCode.Params[match[2]] = match[3]
		}

		result.CQcodes = append(result.CQcodes, cqCode)

		// 替换掉当前CQ码句段
		result.Text = strings.Replace(result.Text, match[0], "", 1)
	}

	// 去除文本中的多余空格
	result.Text = strings.TrimSpace(result.Text)

	return result
}

// 生成CQ码字符串
func GenerateCQCode(cq CQCode) string {
	cqCode := fmt.Sprintf("[CQ:%s", cq.Type)

	for key, value := range cq.Params {
		cqCode += fmt.Sprintf(",%s=%s", key, value)
	}

	return cqCode
}
