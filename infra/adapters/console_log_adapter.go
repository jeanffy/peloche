package adapters

import (
	"log"
	"peloche/domain/ports"
	"time"
)

type LogAdapter struct {
}

func NewLogAdapter() *LogAdapter {
	return &LogAdapter{}
}

func (x *LogAdapter) Info(params ports.LogPortParams) {
	log.Printf("[%s] INF %s: %s\n", time.Now().UTC().String(), params.Module, params.Msg)
}

func (x *LogAdapter) Error(params ports.LogPortErrorParams) {
	log.Printf("[%s] ERR %s: %s - %s\n", time.Now().UTC().String(), params.Module, params.Error.Error(), params.Msg)
}
