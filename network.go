package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PostUrl(body MsgBody) {
	marshal, err := json.Marshal(body)
	if err != nil {
		println(err)
		return
	}
	fmt.Printf("%s", marshal)
	post, err := http.Post(ConfigData.HookURL, "application/json", strings.NewReader(string(marshal)))
	if err != nil {
		println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(post.Body)

	all, err := io.ReadAll(post.Body)
	if err != nil {
		println(err)
		return
	}

	resMsg := &responseMsg{}
	err = json.Unmarshal(all, resMsg)
	if err != nil || resMsg.StatusCode != 0 {
		println("发送失败")
		return
	}
}
