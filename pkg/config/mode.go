package config

import (
	"discord/pkg/logger"
	"fmt"
	"log"
	"os"
)

func General() string {
	loglevel := os.Getenv("log")
	log.Printf("loglevel: %s", loglevel)

	logger.Load("loglevel")
	mode := os.Getenv("mode")
	if mode != "webhook" && mode != "rabbit" {
		logger.Level("fatal", "General", fmt.Sprintf("mode [%s] Not FOund", mode))
	}

	logger.Trace("", " ========= General =========")
	logger.Trace("mode   : ", mode)

	return mode
}
