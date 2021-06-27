package config

import (
	"gopkg.in/ini.v1"
	"log"
)

type Liquid struct {
	ApiTokenId     string
	ApiTokenSecret string
	CancelId       int
	ChangeId       int
	ChangeQuantity string
	ChangePrice    string
}

type Example struct {
	Liquid Liquid
}

func GenerateConfigExample() Example {
	cfg, err := ini.Load("./liquid/config/config.ini")
	if err != nil {
		log.Fatal("read config error", err.Error())
	}
	cancelId, err := cfg.Section("liquid").Key("cancel_id").Int()
	if err != nil {
		log.Fatal("cancel id", err.Error())
	}
	changeId, err := cfg.Section("liquid").Key("change_id").Int()
	if err != nil {
		log.Fatal("cancel id", err.Error())
	}
	return Example{
		Liquid: Liquid{
			ApiTokenId:     cfg.Section("liquid").Key("api_token_id").String(),
			ApiTokenSecret: cfg.Section("liquid").Key("api_token_secret").String(),
			CancelId:       cancelId,
			ChangeId:       changeId,
			ChangeQuantity: cfg.Section("liquid").Key("change_quantity").String(),
			ChangePrice:    cfg.Section("liquid").Key("change_price").String(),
		},
	}
}
