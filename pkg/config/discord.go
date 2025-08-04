package config

import (
	"discord/pkg/logger"
	"os"
)

type settingDiscord struct {
	Token   string
	Channel map[string]string
}

func MyDiscord() settingDiscord {
	//for filter/send
	var channelID = map[string]string{
		os.Getenv("channel_name"): os.Getenv("channel_id"),
	}

	setting := settingDiscord{
		Token:   os.Getenv("token"),
		Channel: channelID,
	}

	logger.Trace("", " ========= MyDiscord =========")
	logger.Trace("Token  : ", setting.Token)
	logger.Trace("Channel: ", channelID)
	return setting
}
