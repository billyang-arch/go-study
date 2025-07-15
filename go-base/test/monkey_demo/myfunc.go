package monkey_demo

import "fmt"

// func.go

func MyFunc(uid int64) string {
	u, err := GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// 这里是一些逻辑代码...

	return fmt.Sprintf("hello %s\n", u)
}

func GetInfoByUID(uid int64) (u string, err error) {
	return "", nil
}
