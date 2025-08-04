package main

import (
	"context"
	"discord/pkg/config"
	"discord/pkg/discord"
	"discord/pkg/logger"
	"discord/pkg/rabbit"
	"discord/pkg/webhook"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	mode := config.General()
	settingDiscord := config.MyDiscord()
	settingWebHook := config.Webhook()
	settingRabbit := config.Rabbit()

	repoWH := webhook.NewRepo(settingWebHook)
	repoRabbit := rabbit.NewRabbit(settingRabbit)
	if mode == "rabbit" {
		if err := repoRabbit.Connect(ctx); err != nil {
			logger.Level("error", "main", fmt.Sprintf("repoRabbit.Connect:%v", err))
			cancel()
		}
		logger.Level("info", "Start", "succes connect rabbit")
	}

	//discord
	_, err := discord.NewRepo(settingDiscord.Token, settingDiscord.Channel, mode, repoWH, repoRabbit)
	if err != nil {
		logger.Level("error", "main", fmt.Sprintf("discord.NewRepo:%v", err))
		cancel()
	}
	logger.Level("info", "Start", "succes connect discord")
	time.Sleep(1 * time.Second)

	log.Println("===== Run  =====")
	<-ctx.Done()
}
