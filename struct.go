package main

type responseMsg struct {
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
}

type MsgBody struct {
	Timestamp string      `json:"timestamp,omitempty"`
	Sign      string      `json:"sign,omitempty"`
	MsgType   msgType     `json:"msg_type,omitempty"`
	Content   interface{} `json:"content,omitempty"`
}

type MsgContentText struct {
	Text string `json:"text,omitempty"`
}

type MsgContentShareChatId struct {
	ShareChatId string `json:"share_chat_id,omitempty"`
}

type MsgContentImageKey struct {
	ImageKey string `json:"image_key,omitempty"`
}

type MsgContentPost struct {
	Post MsgContentPostPost `json:"post,omitempty"`
}

type MsgContentPostPost struct {
	ZhCn MsgContentPostPostTemplate `json:"zh_cn,omitempty"`
	EnUs MsgContentPostPostTemplate `json:"en_us,omitempty"`
}

type MsgContentPostPostTemplate struct {
	Title   string                                `json:"title,omitempty"`
	Content [][]MsgContentPostPostTemplateContent `json:"content,omitempty"`
}

type MsgContentPostPostTemplateContent struct {
	Tag  postContentTag `json:"tag,omitempty"`
	Text struct {
		Text     string `json:"text,omitempty"`
		UnEscape bool   `json:"un_escape,omitempty"`
	}
	A struct {
		Text string `json:"text,omitempty"`
		Href string `json:"href,omitempty"`
	}
	At struct {
		UserId   string `json:"user_id,omitempty"`
		UserName string `json:"user_name,omitempty"`
	}
	Img struct {
		ImageKey string `json:"image_key,omitempty"`
	}
	Media struct {
		FileKey  string `json:"file_key,omitempty"`
		ImageKey string `json:"image_key,omitempty"`
	}
	Emotion struct {
		EmojiType string `json:"emoji_type,omitempty"`
	}
}

// MsgContentInteractive Card内容通过以下地址生成：
//
// https://open.feishu.cn/tool/cardbuilder?from=custom_bot_doc
type MsgContentInteractive struct {
	Card string `json:"card,omitempty"`
}
