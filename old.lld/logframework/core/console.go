package core

import (
	"fmt"
	log "lld/logframework/logger"
)

type ConsoleAppender struct {
}

func (a *ConsoleAppender) Name() string {
	return "ConsoleAppender"
}

func (a *ConsoleAppender) Append(message *log.LogMessage) error {
	fmt.Println(message)
	return nil
}
