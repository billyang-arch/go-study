// client/main.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type addParam1 struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResult1 struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

// RESTful API多用于前后端之间的数据传输，而目前微服务架构下各个微服务之间多采用RPC调用。
func main() {
	// 通过HTTP请求调用其他服务器上的add服务
	url := "http://127.0.0.1:9090/add"
	param := addParam1{
		X: 10,
		Y: 20,
	}
	paramBytes, _ := json.Marshal(param)
	resp, _ := http.Post(url, "application/json", bytes.NewReader(paramBytes))
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	var respData addResult1
	json.Unmarshal(respBytes, &respData)
	fmt.Println(respData.Data) // 30
}
