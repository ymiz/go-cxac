package config

import (
	"gopkg.in/ini.v1"
	"log"
)

type Bitflyer struct {
	AccessKey string
	SecretKey string
	CancelId  string
}

type Example struct {
	Bitflyer Bitflyer
}

func GenerateConfigExample() Example {
	cfg, err := ini.Load("./bitflyer/config/config.ini")
	if err != nil {
		log.Fatal("read config error", err.Error())
	}
	return Example{
		Bitflyer: Bitflyer{
			AccessKey: cfg.Section("bitflyer").Key("access_key").String(),
			SecretKey: cfg.Section("bitflyer").Key("secret_key").String(),
			CancelId:  cfg.Section("bitflyer").Key("cancel_id").String(),
		},
	}
}
