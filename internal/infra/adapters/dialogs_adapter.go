package adapters

import (
	"peloche/internal/infra/ui"
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	sdialog "github.com/sqweek/dialog"
)

// ---------------------------------------------------------------------------
// #region definition

var _ ui.DialogsPort = (*DialogsAdapter)(nil)

type DialogsAdapter struct {
	fyneApp fyne.App
	router  ui.RouterPort
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewDialogsAdapter() *DialogsAdapter {
	return &DialogsAdapter{
		fyneApp: di.GetBasicDI().Resolve("FyneApp").(fyne.App),
		router:  di.GetBasicDI().Resolve(ui.ROUTER_PORT_TOKEN).(ui.RouterPort),
	}
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *DialogsAdapter) MessageDialog(msg string) {
	sdialog.Message("%s", msg).Info()
	// FIXME: when dialog is closed, parent window does not get the focus back
}

func (x *DialogsAdapter) ErrorDialog(err error) {
	parent := x.router.GetCurrentWindow()
	dialog.NewError(err, parent).Show()
}

// #endregion

// ---------------------------------------------------------------------------
// #region events

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion
