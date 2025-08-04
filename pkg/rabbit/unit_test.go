package rabbit_test

import (
	"context"
	"discord/pkg/rabbit"
	"encoding/json"
	"testing"
)

var setting = rabbit.Setting{
	Host: "amqp://user:pass@localhost:5672/",
	Tag:  "n8n-discord",
	Que:  "n8n",
	RoutingKey: map[string]string{
		"ask-me": "n8n.discord.ask-me",
	},
}

func Test_Publishx1B(t *testing.T) {
	ctx := context.Background()
	repo := rabbit.NewRabbit(setting)
	if err := repo.Connect(ctx); err != nil {
		t.Fatalf("Connect:%v", err)
	}

	payload := map[string]interface{}{
		"channelID": "123",
		"messageID": "456",
		"content":   "apa node discord tersedia ?",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Marshal:%v", err)
	}

	if err := repo.Publish( "ask-me", body); err != nil {
		t.Fatalf("Publish:%v", err)
	}
}
