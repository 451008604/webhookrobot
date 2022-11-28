package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	HookURL string `json:"hookURL"`
	SignKey string `json:"signKey"`
}

var ConfigData = &Config{}

func init() {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
		return
	}
	println(string(file))
	err = json.Unmarshal(file, ConfigData)
	if err != nil {
		panic(err)
		return
	}
}
