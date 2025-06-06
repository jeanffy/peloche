package ui

import (
	"peloche/infra/ui/routing"

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

func NewUIDialogs(fyneApp fyne.App, router routing.Router) *UIDialogs {
	return &UIDialogs{
		fyneApp: fyneApp,
		router:  router,
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
