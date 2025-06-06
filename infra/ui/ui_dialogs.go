package ui

import (
	"peloche/infra/ui/routing"
	"peloche/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	sdialog "github.com/sqweek/dialog"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type UIDialogs struct {
	fyneApp fyne.App
	router  routing.Router
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewUIDialogs() *UIDialogs {
	return &UIDialogs{
		fyneApp: utils.GetNaiveDI().Resolve("FyneApp").(fyne.App),
		router:  utils.GetNaiveDI().Resolve(routing.ROUTER_TOKEN).(routing.Router),
	}
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *UIDialogs) MessageDialog(msg string) {
	sdialog.Message("%s", msg).Info()
	// FIXME: when dialog is closed, parent window does not get the focus back
}

func (x *UIDialogs) ErrorDialog(err error) {
	parent := x.router.GetCurrentWindow()
	dialog.NewError(err, parent).Show()
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------
