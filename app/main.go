package main

import (
	// json

	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 定义结构体
type interview struct {
	Message string
	// Timestamp string
	Timestamp int
}

// 程序入口
func main() {
	request()
}

// 自定义func
func tadyinterview(response http.ResponseWriter, r *http.Request) {
	interviewtype := []interview{
		interview{
			Message:   "tady interview",
			Timestamp: 07022021,
		},
	}

	json.NewEncoder(response).Encode(interviewtype)

	fmt.Println("Endpoint", interviewtype)
}

func tadyinterview_test(response http.ResponseWriter, r *http.Request) {
	interviewtype := []interview{
		interview{
			Message:   "hello world /tadyinterview_test",
			Timestamp: 07022021,
		},
	}

	json.NewEncoder(response).Encode(interviewtype)

	fmt.Println("Endpoint", interviewtype)
}

// http请求
func request() {
	// 处理func，url转发
	http.HandleFunc("/", tadyinterview)

	http.HandleFunc("/test", tadyinterview_test)

	log.Fatal(http.ListenAndServe(":9999", nil))
}
