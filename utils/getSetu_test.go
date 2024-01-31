package utils

import (
	"testing"
)

func Test_getsetu(t *testing.T) {
	// 调用 Get_setu 函数
	result := Get_setu([]string{"白丝"}, 1, 1)
	t.Logf(result.Data[0].Urls.Regular)
}
