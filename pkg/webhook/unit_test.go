package webhook_test

import (
	"discord/pkg/webhook"
	"testing"
)

var setting = webhook.Setting{
	Url:   "https://ip/webhook/discord",
	Name:  "AUTH-NAME",
	Value: "AUTH-VALUE",
}

func Test_Postx(t *testing.T) {
	payload := map[string]interface{}{
		"channelID": "123",
		"messageID": "456",
		"content":   "n8n itu apa ?",
	}

	repo := webhook.NewRepo(setting)
	if err := repo.POSTAskMe(payload); err != nil {
		t.Fatalf("POSTAskMe:%v", err)
	}
}
