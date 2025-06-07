package adapters

import (
	"log"
	"peloche/internal/domain"
	"time"
)

type ConsoleLogAdapter struct {
}

func NewConsoleLogAdapter() *ConsoleLogAdapter {
	return &ConsoleLogAdapter{}
}

func (x *ConsoleLogAdapter) Info(params domain.LogPortParams) {
	log.Printf("[%s] INF %s: %s\n", time.Now().UTC().String(), params.Module, params.Msg)
}

func (x *ConsoleLogAdapter) Error(params domain.LogPortErrorParams) {
	log.Printf("[%s] ERR %s: %s - %s\n", time.Now().UTC().String(), params.Module, params.Error.Error(), params.Msg)
}
