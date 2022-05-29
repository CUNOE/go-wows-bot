package global

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Token           string `json:"token"`
	ServerStartPort string `json:"server_start_port"`
}

func LoadConfig() {
	var conf Config

	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("缺少config.json文件", err.Error())

	}

	err = json.Unmarshal(jsonFile, &conf)
	if err != nil {
		log.Fatal("解析错误 请检查你的json语法", err.Error())

	}

	ServerStartPort = conf.ServerStartPort
	Token = conf.Token
}
