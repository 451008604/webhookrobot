package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

func GenerateSign(secret string) (sign string, timestamp int64, err error) {
	timestamp = time.Now().Unix()
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err = h.Write(data)
	if err != nil {
		return "", 0, err
	}

	sign = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}
