package adapters

import (
	"log"
	"peloche/internal/domain"
	"time"
)

type LogAdapter struct {
}

func NewLogAdapter() *LogAdapter {
	return &LogAdapter{}
}

func (x *LogAdapter) Info(params domain.LogPortParams) {
	log.Printf("[%s] INF %s: %s\n", time.Now().UTC().String(), params.Module, params.Msg)
}

func (x *LogAdapter) Error(params domain.LogPortErrorParams) {
	log.Printf("[%s] ERR %s: %s - %s\n", time.Now().UTC().String(), params.Module, params.Error.Error(), params.Msg)
}
