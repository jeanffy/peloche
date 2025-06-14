package adapters

import (
	"log"
	"peloche/internal/domain"
	"time"
)

// ---------------------------------------------------------------------------
// #region definition

var _ domain.LogPort = (*ConsoleLogAdapter)(nil)

type ConsoleLogAdapter struct {
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewConsoleLogAdapter() *ConsoleLogAdapter {
	return &ConsoleLogAdapter{}
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *ConsoleLogAdapter) Info(params domain.LogPortParams) {
	log.Printf("[%s] INF %s: %s\n", time.Now().UTC().String(), params.Module, params.Msg)
}

func (x *ConsoleLogAdapter) Error(params domain.LogPortErrorParams) {
	log.Printf("[%s] ERR %s: %s - %s\n", time.Now().UTC().String(), params.Module, params.Error.Error(), params.Msg)
}

// #endregion

// ---------------------------------------------------------------------------
// #region events

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion
