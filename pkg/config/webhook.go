package config

import (
	"discord/pkg/logger"
	"discord/pkg/webhook"
	"os"
)

func Webhook() webhook.Setting {
	setting := webhook.Setting{
		Url:   os.Getenv("url"),
		Name:  os.Getenv("name"),
		Value: os.Getenv("value"),
	}
	logger.Trace("", " ========= Webhook =========")
	logger.Trace("Url    : ", setting.Url)
	logger.Trace("Name   : ", setting.Name)
	logger.Trace("Value  : ", setting.Value)

	return setting
}
