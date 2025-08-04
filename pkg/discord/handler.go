package discord

import (
	"encoding/json"
	"fmt"
	"strings"

	"discord/pkg/logger"

	"github.com/bwmarrin/discordgo"
)

type Payload struct {
	ChannelID string `json:"channelID"`
	MessageID string `json:"messageID"`
	Content   string `json:"content"`
}

func (r *repo) handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	logger.Level("debug", "handler", fmt.Sprintf("ChannelID: %s", m.ChannelID))
	logger.Level("debug", "handler", fmt.Sprintf("Author   : %s", m.Author.ID))

	if m.Author.ID == s.State.User.ID {
		return
	}

	logger.Level("debug", "handler", fmt.Sprintf("MessageID: %s", m.ID))
	logger.Level("debug", "handler", fmt.Sprintf("Content  : %s", m.Content))

	for alias, id := range r.channelID {
		if m.ChannelID != id {
			continue
		}
		logger.Level("debug", "handler", fmt.Sprintf("channel match: %s", alias))
		//switch case channel
		if alias != "ask-me" {
			continue
		}

		if !strings.Contains(m.Content, "#sahabot") {
			continue
		}
		logger.Level("debug", "handler", "tag #sahabo found")

		content := strings.ReplaceAll(m.Content, "#sahabot", "")
		payload := map[string]interface{}{
			"channelID": m.ChannelID,
			"messageID": m.ID,
			"content":   content,
		}

		body, err := json.Marshal(payload)
		if err != nil {
			logger.Level("error", "handler", fmt.Sprintf("Marshal:%v", err))
			return
		}
		logger.Level("debug", "handler", fmt.Sprintf("Payload: \n%s", string(body)))

		if r.mode == "webhook" {
			if err := r.repoWH.POSTAskMe(payload); err != nil {
				logger.Level("error", "handler", fmt.Sprintf("POSTAskMe:%v", err))
				return
			}
			logger.Level("debug", "handler", "repoWH.POSTAskMe succes")
		} else if r.mode == "rabbit" {
			if err := r.repoRabbit.Publish(alias, body); err != nil {
				logger.Level("error", "handler", fmt.Sprintf("Publish:%v", err))
				return
			}
			logger.Level("debug", "handler", "repoRabbit.Publish succes")
		}

	}
}
