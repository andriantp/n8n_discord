package config

import (
	"discord/pkg/logger"
	"discord/pkg/rabbit"
	"os"
)

func Rabbit() rabbit.Setting {
	setting := rabbit.Setting{
		Host: os.Getenv("host"),
		Tag:  os.Getenv("tag"),
		Que:  os.Getenv("que"),
		RoutingKey: map[string]string{
			os.Getenv("channel_name"): os.Getenv("routing_key"),
		},
	}
	logger.Trace("", " ========= Rabbit =========")
	logger.Trace("Host   : ", setting.Host)
	logger.Trace("Tag    : ", setting.Tag)
	logger.Trace("Que    : ", setting.Que)
	logger.Trace("RK     : ", setting.RoutingKey)
	return setting
}
