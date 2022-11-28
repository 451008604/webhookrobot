package main

func main() {
	body := getMsgBody(Text)
	body.Content = MsgContentText{Text: "666"}
	PostUrl(body)
}
