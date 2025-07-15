package monkey_demo

import (
	"bou.ke/monkey"
	"strings"
	"testing"
)

// func_test.go
// go test -run=TestMyFunc -v -gcflags=-l 防止内联优化
// monkey原理是使用汇编语言重写可执行文件，将目标函数或方法的实现跳转到桩实现，其原理类似于热补丁。
// monkey不支持内联函数，在测试的时候需要通过命令行参数-gcflags=-l关闭Go语言的内联优化。
// monkey不是线程安全的，所以不要把它用到并发的单元测试中。
func TestMyFunc(t *testing.T) {
	// 对 GetInfoByUID 进行打桩
	// 无论传入的uid是多少，都返回 &varys.UserInfo{Name: "liwenzhou"}, nil
	monkey.Patch(GetInfoByUID, func(int64) (string, error) {
		return "nihao", nil
	})

	ret := MyFunc(123)
	if !strings.Contains(ret, "nihao") {
		t.Fatal()
	}
}
