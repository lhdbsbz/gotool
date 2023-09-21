package gotool

import (
	"encoding/json"
	"net/http"
	"strings"
)

func SendToFlyBook(url, content string) {
	if len(url) == 0 || len(content) == 0 {
		return
	}
	body := struct {
		MsgType string `json:"msg_type"`
		Content struct {
			Text string `json:"text"`
		} `json:"content"`
	}{
		MsgType: "text",
		Content: struct {
			Text string `json:"text"`
		}{
			Text: content,
		},
	}
	bodyStr, _ := json.Marshal(body)
	_, _ = http.Post(url, "application/json;charset=utf-8", strings.NewReader(string(bodyStr)))
}
