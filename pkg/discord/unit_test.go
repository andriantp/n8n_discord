package discord_test

import (
	"discord/pkg/discord"
	"discord/pkg/rabbit"
	"discord/pkg/webhook"
	"testing"
)

var (
	token     = ""
	channelID = map[string]string{
		"channel": "13",
	}
	messageID = "13"
	image     = "image/image.png"

	settingWebHook = webhook.Setting{}
	settingRabbit  = rabbit.Setting{}
)

func Test_SendText(t *testing.T) {
	repoWH := webhook.NewRepo(settingWebHook)
	repoRabbit := rabbit.NewRabbit(settingRabbit)
	repo, err := discord.NewRepo(token, channelID, "", repoWH, repoRabbit)
	if err != nil {
		t.Fatalf("NewRepo: %v", err)
	}

	if err := repo.SendText(channelID["channel"], "#sahabot apakah node discord tersedia onMessage ?"); err != nil {
		t.Fatalf("SendText: %v", err)
	}
}
