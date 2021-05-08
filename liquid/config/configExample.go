package config

import (
	"gopkg.in/ini.v1"
	"log"
)

type Liquid struct {
	ApiTokenId     string
	ApiTokenSecret string
}

type Example struct {
	Liquid Liquid
}

func GenerateConfigExample() Example {
	cfg, err := ini.Load("./liquid/config/config.ini")
	if err != nil {
		log.Fatal("read config error", err.Error())
	}
	return Example{
		Liquid: Liquid{
			ApiTokenId:     cfg.Section("liquid").Key("api_token_id").String(),
			ApiTokenSecret: cfg.Section("liquid").Key("api_token_secret").String(),
		},
	}
}
