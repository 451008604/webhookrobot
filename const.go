package main

type msgType string

const (
	Text        msgType = "text"
	Post        msgType = "post"
	Image       msgType = "image"
	ShareChat   msgType = "share_chat"
	Interactive msgType = "interactive"
)

type postContentTag string

const (
	TagText    postContentTag = "text"
	TagA       postContentTag = "a"
	TagAt      postContentTag = "at"
	TagImg     postContentTag = "img"
	TagMedia   postContentTag = "media"
	TagEmotion postContentTag = "emotion"
)
