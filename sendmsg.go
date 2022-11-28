package main

import "strconv"

func getMsgBody(msgType msgType) MsgBody {
	sign, timestamp, err := GenerateSign(ConfigData.SignKey)
	if err != nil {
		return MsgBody{}
	}
	body := MsgBody{}
	body.Timestamp = strconv.FormatInt(timestamp, 10)
	body.Sign = sign
	body.MsgType = msgType
	return body
}

// SendText 发送文本消息
func SendText(data MsgContentText) {
	body := getMsgBody(Text)
	body.Content = data
	PostUrl(body)
}

// SendPost 发送富文本消息
func SendPost(data MsgContentPost) {
	body := getMsgBody(Post)
	body.Content = data
	PostUrl(body)
}

// SendImage 发送图片消息
func SendImage(data MsgContentImageKey) {
	body := getMsgBody(Image)
	body.Content = data
	PostUrl(body)
}

// SendShareChat 发送群名片消息
func SendShareChat(data MsgContentShareChatId) {
	body := getMsgBody(ShareChat)
	body.Content = data
	PostUrl(body)
}

// SendInteractive 发送消息卡片
func SendInteractive(data MsgContentInteractive) {
	body := getMsgBody(Interactive)
	body.Content = data
	PostUrl(body)
}
